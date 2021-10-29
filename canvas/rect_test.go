package canvas

import "testing"

func TestRectEndCoord(t *testing.T) {
	rect := NewRectangle(
		[2]int{1, 2},
		2,
		2,
		"x",
		"o",
	)

	endX, endY := rect.EndCoord()
	if endX != 2 {
		t.Errorf("expected end x as 2 got %d", endX)
	}
	if endY != 3 {
		t.Errorf("expected end y as 3 got %d", endY)
	}
}

func TestPixel(t *testing.T) {
	testCases := []struct {
		name     string
		rect     *Rectangle
		point    [2]int
		expected string
	}{
		{
			name:     "NormalCase",
			rect:     NewRectangle([2]int{0, 0}, 3, 3, "x", "o"),
			point:    [2]int{1, 1},
			expected: "o",
		},
		{
			name:     "IsOutline",
			rect:     NewRectangle([2]int{1, 3}, 3, 2, "x", "o"),
			point:    [2]int{1, 4},
			expected: "x",
		},
		{
			name:     "NoOutline",
			rect:     NewRectangle([2]int{1, 3}, 3, 3, "", "."),
			point:    [2]int{1, 4},
			expected: ".",
		},
		{
			name:     "NoFill",
			rect:     NewRectangle([2]int{1, 1}, 5, 5, "o", ""),
			point:    [2]int{2, 2},
			expected: " ",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			gotten := test.rect.Pixel(test.point[0], test.point[1])
			if gotten != test.expected {
				t.Errorf("expected '%s' got '%s'", test.expected, gotten)
			}
		})
	}
}
