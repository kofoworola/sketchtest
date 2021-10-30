package main

import (
	"context"
	"net/http"

	"github.com/kofoworola/sketchtest/storage"
)

type Handler struct {
	store canvasStorage
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/canvas" {
		h.drawHandler(w, req)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// use an interface in case of mock storage
type canvasStorage interface {
	CreateOrUpdateCanvas(ctx context.Context, input storage.Canvas) (*storage.Canvas, error)
	GetCanvasById(ctx context.Context, id string) (*storage.Canvas, error)
}

func respond(body string, code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	w.Write([]byte(body))
}
