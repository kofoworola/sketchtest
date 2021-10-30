package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kofoworola/sketchtest/canvas"
)

func (h *Handler) drawHandler(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid method")
		return
	}

	if err := req.ParseForm(); err != nil {
		log.Printf("error: %v", err)
		respond("", http.StatusBadRequest, w)
		return
	}
	id := req.Form.Get("id")
	if id == "" {
		respond("id is required", http.StatusBadRequest, w)
		return
	}

	var canv *canvasReq
	// get request body to check if it is empty
	body, err := io.ReadAll(req.Body)
	if len(body) < 1 {
		canv, err = h.getCanvasFromStorage(req.Context(), id)
		if err != nil {
			if err == sql.ErrNoRows {
				respond("canvas not found, create one by passing body", http.StatusNotFound, w)
				return
			}
			log.Printf("error: %v", err)
			respond("error getting canvas", http.StatusInternalServerError, w)
			return
		}
	} else {
		canv, err = h.getCanvasFromBody(body)
		if err != nil {
			respond(err.Error(), http.StatusBadRequest, w)
			return
		}
	}

	draw(canv, w)

	// TODO implement retries/queue to storage service
	if _, err := h.store.CreateOrUpdateCanvas(req.Context(), canv.toStorage(id)); err != nil {
		log.Printf("error creating canvas in db: %v", err)
	}

}

// getCanvasFromBody gets the canvas data from the passed body and validates it
func (h *Handler) getCanvasFromBody(body []byte) (*canvasReq, error) {
	var canvas canvasReq
	if err := json.Unmarshal(body, &canvas); err != nil {
		return nil, errors.New("invalid body")
	}
	if err := validate(&canvas); err != nil {
		return nil, err
	}
	return &canvas, nil
}

// getCanvasFromStorage gets the canvas data from the postgres db
func (h *Handler) getCanvasFromStorage(ctx context.Context, id string) (*canvasReq, error) {
	canvas, err := h.store.GetCanvasById(ctx, id)
	if err != nil {
		return nil, err
	}
	rectangles := make([]rectangle, len(canvas.Operations.Rectangles))
	fills := make([]fill, len(canvas.Operations.Fills))

	for i, rect := range canvas.Operations.Rectangles {
		rectangles[i] = rectangle{
			StartX:  rect.StartX,
			StartY:  rect.StartY,
			Height:  rect.Height,
			Fill:    rect.Fill,
			Outline: rect.Outline,
			Width:   rect.Width,
		}
	}

	for i, f := range canvas.Operations.Fills {
		fills[i] = fill{
			StartX:    f.StartX,
			StartY:    f.StartY,
			Character: f.Character,
		}
	}

	return &canvasReq{
		Rectangles: rectangles,
		Fills:      fills,
	}, nil
}

func draw(input *canvasReq, w http.ResponseWriter) {
	operations := make([]canvas.Operation, 0)
	for _, f := range input.Fills {
		operations = append(operations, canvas.NewFill([2]int{f.StartX, f.StartY}, f.Character))
	}
	for _, rect := range input.Rectangles {
		operations = append(operations, canvas.NewRectangle(
			[2]int{rect.StartX, rect.StartY},
			rect.Width,
			rect.Height,
			rect.Outline,
			rect.Fill,
		))
	}

	canvas := canvas.NewCanvas(operations...)
	canvas.Draw(w)
}

func validate(input *canvasReq) error {
	return validation.ValidateStruct(
		input,
		validation.Field(&input.Rectangles, validation.Required, validation.By(validateRectArray)),
		validation.Field(&input.Fills, validation.By(validateFillArray)),
	)
}

func validateRectArray(a interface{}) error {
	rects, ok := a.([]rectangle)
	if !ok {
		validation.NewError("invalid type", "rectangles should be an array of rectangles")
	}
	for _, item := range rects {
		if err := validation.ValidateStruct(
			&item,
			validation.Field(&item.StartX, validation.Min(0)),
			validation.Field(&item.StartY, validation.Min(0)),
			validation.Field(&item.Width, validation.Required, validation.Min(0)),
			validation.Field(&item.Height, validation.Required, validation.Min(0)),
			validation.Field(&item.Fill, validation.When(item.Outline == "", validation.Required), validation.Length(0, 1)),
			validation.Field(&item.Outline, validation.When(item.Fill == "", validation.Required), validation.Length(0, 1)),
		); err != nil {
			return err
		}
	}
	return nil
}

func validateFillArray(a interface{}) error {
	fills, ok := a.([]fill)
	if !ok {
		validation.NewError("invalid_type", "fills should be an array of fills")
	}
	for _, item := range fills {
		if err := validation.ValidateStruct(
			&item,
			validation.Field(&item.StartX, validation.Min(0)),
			validation.Field(&item.StartY, validation.Min(0)),
			validation.Field(&item.Character, validation.Required, validation.Length(0, 1)),
		); err != nil {
			return err
		}
	}
	return nil
}
