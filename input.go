package main

import "github.com/kofoworola/sketchtest/storage"

type canvasReq struct {
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

// toStorage converts canvasReq to a storage.Canvas
func (c *canvasReq) toStorage(id string) storage.Canvas {
	rectangles := make([]storage.Rectangle, len(c.Rectangles))
	fills := make([]storage.Fill, len(c.Fills))

	for i, rect := range c.Rectangles {
		rectangles[i] = storage.Rectangle{
			StartX:  rect.StartX,
			StartY:  rect.StartY,
			Fill:    rect.Fill,
			Outline: rect.Outline,
			Height:  rect.Height,
			Width:   rect.Width,
		}
	}

	for i, fill := range c.Fills {
		fills[i] = storage.Fill{
			StartX:    fill.StartX,
			StartY:    fill.StartY,
			Character: fill.Character,
		}
	}

	return storage.Canvas{
		ID: id,
		Operations: storage.Operations{
			Rectangles: rectangles,
			Fills:      fills,
		},
	}
}
