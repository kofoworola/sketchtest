package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kofoworola/sketchtest/canvas"
	"github.com/kofoworola/sketchtest/storage"
)

type Handler struct {
	store canvasStorage
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.RequestURI == "/canvas" {
		h.drawHandler(w, req)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// use an interface in case of mock storage
type canvasStorage interface {
	CreateCanvas(ctx context.Context, input storage.Canvas) (*storage.Canvas, error)
}

type inputBody struct {
	Rectangles []rectangle `json:"rectangles"`
	Fills      []fill      `json:"fills"`
}

type rectangle struct {
	StartX  int    `json:"start_x"`
	StartY  int    `json:"start_y"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Outline string `json:"outline"`
	Fill    string `json:"fill"`
}

type fill struct {
	StartX    int    `json:"start_x"`
	StartY    int    `json:"start_y"`
	Character string `json:"character"`
}

func (h *Handler) drawHandler(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid method")
		return
	}

	// input and validation
	var input inputBody
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := validate(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// operation
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
	// parse values
}

func validate(input *inputBody) error {
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
