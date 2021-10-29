package canvas

import (
	"io"
	"strings"
)

type operationType string

const (
	typeRectangle operationType = "rectangle"
	typeFill      operationType = "fill"
)

type Canvas struct {
	operations []operation

	endX int
	endY int
}

type operation interface {
	Type() operationType
	Pixel(int, int) string
	EndCoord() (int, int)
}

func NewCanvas(operations ...operation) *Canvas {
	canvas := Canvas{operations: operations}

	// find endX and endY
	var endX, endY = 0, 0
	for _, op := range operations {
		if op.Type() == typeFill {
			continue
		}
		x, y := op.EndCoord()
		if x > endX {
			endX = x
		}
		if y > endY {
			endY = y
		}
	}
	// add extra space to end coord
	canvas.endX = endX + 1
	canvas.endY = endY + 1

	return &canvas
}

func (c *Canvas) Draw(writer io.Writer) error {
	var builder strings.Builder
	for y := 0; y <= c.endY; y++ {
		for x := 0; x <= c.endX; x++ {
			val := " "
			for _, r := range c.operations {
				pixel := r.Pixel(x, y)
				if pixel == "" {
					continue
				}
				val = pixel
			}
			builder.WriteString(val)
		}
		builder.WriteString("\n")
	}
	if _, err := writer.Write([]byte(builder.String())); err != nil {
		return err
	}
	return nil
}
