package canvas

import (
	"testing"
)

func TestCanvasInit(t *testing.T) {
	rects := []Operation{
		NewRectangle([2]int{1, 1}, 3, 3, "x", "o"),
		NewRectangle([2]int{1, 2}, 2, 4, "x", "o"),
	}

	canvas := NewCanvas(rects...)
	if canvas.endX != 4 {
		t.Errorf("expected 3 got %d", canvas.endX)
	}
	if canvas.endY != 6 {
		t.Errorf("expected 5 got %d", canvas.endY)
	}
}
