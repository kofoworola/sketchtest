package canvas

type Fill struct {
	startX        int
	startY        int
	fillCharacter string
}

func NewFill(start [2]int, character string) *Fill {
	return &Fill{
		startX:        start[0],
		startY:        start[1],
		fillCharacter: character,
	}
}

func (f *Fill) Type() operationType { return typeFill }

func (f *Fill) EndCoord() (x, y int) { return 0, 0 }

func (f *Fill) Pixel(x, y int) string {
	if x >= f.startX && y >= f.startY {
		return f.fillCharacter
	}
	return ""
}
