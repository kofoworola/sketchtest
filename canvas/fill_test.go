package canvas

import "testing"

func TestFill(t *testing.T) {
	testCases := []struct {
		name     string
		fill     *Fill
		coords   [2]int
		expected string
	}{
		{
			name:     "NormalFill",
			fill:     NewFill([2]int{0, 0}, "."),
			coords:   [2]int{5, 5},
			expected: ".",
		},
		{
			name:     "OutOfBounds",
			fill:     NewFill([2]int{1, 3}, "."),
			coords:   [2]int{0, 2},
			expected: "",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			gotten := test.fill.Pixel(test.coords[0], test.coords[1])
			if gotten != test.expected {
				t.Errorf("expected %s got %s", test.expected, gotten)
			}
		})
	}
}
