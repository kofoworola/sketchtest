package canvas

// todo move to canvas
type Rectangle struct {
	StartX int
	StartY int

	outline string
	fill    string
	width   int
	height  int
	endX    int
	endY    int
}

func NewRectangle(start [2]int, width, height int, outline, fill string) *Rectangle {
	return &Rectangle{
		StartX:  start[0],
		StartY:  start[1],
		outline: outline,
		fill:    fill,
		width:   width,
		height:  height,
		// the end points are the start coords added to the width
		// one is subtracted since the start coord is included
		endX: start[0] + width - 1,
		endY: start[1] + height - 1,
	}
}

func (r *Rectangle) Type() operationType { return typeRectangle }

func (r *Rectangle) EndCoord() (x, y int) { return r.endX, r.endY }

// Pixel checks if the point is in the rectangle and returns the rune to be drawn at that point
func (r *Rectangle) Pixel(x, y int) string {
	// it is inside the rectangle is the x,y is between startX, endX
	// and startY, endY respectively
	insideX := x >= r.StartX && x <= r.endX
	insideY := y >= r.StartY && y <= r.endY
	inside := insideX && insideY
	if !inside {
		return ""
	}

	// it is an outline if x,y is startX, endX and startY, endY respectively
	xOutline := x == r.StartX || x == r.endX
	yOutline := y == r.StartY || y == r.endY
	if xOutline || yOutline {
		if r.outline == "" {
			return r.fill
		}
		return r.outline
	}

	// not an outline
	if r.fill == "" {
		return " "
	}
	return r.fill
}
