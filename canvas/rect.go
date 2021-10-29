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
		endX:    start[0] + width - 1,
		endY:    start[1] + height - 1,
	}
}

func (r *Rectangle) Type() operationType { return typeRectangle }

func (r *Rectangle) EndCoord() (x, y int) { return r.endX, r.endY }

// Pixel checks if the point is in the rectangle and returns the rune to be drawn at that point
func (r *Rectangle) Pixel(x, y int) string {
	insideX := x >= r.StartX && x <= r.endX
	insideY := y >= r.StartY && y <= r.endY
	inside := insideX && insideY
	if !inside {
		return ""
	}

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
