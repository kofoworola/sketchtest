package canvas

import (
	"io"
)

type Canvas struct {
	writer io.Writer

	rects []*Rectangle

	endX int
	endY int
}

func NewCanvas(writer io.Writer, rectangles ...*Rectangle) *Canvas {
	canvas := Canvas{writer: writer, rects: rectangles}

	// find endX and endY
	var endX, endY = 0, 0
	for _, rect := range rectangles {
		if rect.endX > endX {
			endX = rect.endX
		}
		if rect.endY > endY {
			endY = rect.endY
		}
	}
	// add extra space to end coord
	canvas.endX = endX + 1
	canvas.endY = endY + 1

	return &canvas
}
