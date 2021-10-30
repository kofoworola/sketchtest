package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kofoworola/sketchtest/storage"
)

func TestCanvasToStorage(t *testing.T) {
	canv := &canvasReq{
		Rectangles: []rectangle{
			{
				StartX:  0,
				StartY:  0,
				Fill:    "x",
				Outline: "y",
				Height:  5,
				Width:   5,
			},
		},
		Fills: []fill{
			{
				StartX:    0,
				StartY:    0,
				Character: "x",
			},
		},
	}

	expected := storage.Canvas{
		ID: "test",
		Operations: storage.Operations{
			Rectangles: []storage.Rectangle{
				{
					StartX:  0,
					StartY:  0,
					Fill:    "x",
					Outline: "y",
					Height:  5,
					Width:   5,
				},
			},
			Fills: []storage.Fill{
				{
					StartX:    0,
					StartY:    0,
					Character: "x",
				},
			},
		},
	}

	gotten := canv.toStorage("test")
	if diff := cmp.Diff(expected, gotten); diff != "" {
		t.Fatal(diff)
	}
}
