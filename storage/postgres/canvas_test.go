package postgres

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kofoworola/sketchtest/storage"
)

func TestCreateCanvas(t *testing.T) {
	testCases := []struct {
		name   string
		canvas *storage.Canvas
		err    error
	}{
		{
			name: "SuccessfulCreate",
			canvas: &storage.Canvas{
				ID: "test-canvas",
				Operations: storage.Operations{
					Rectangles: []storage.Rectangle{
						{
							StartX:  0,
							StartY:  0,
							Fill:    "o",
							Height:  10,
							Outline: "x",
							Width:   10,
						},
					},
				},
			},
		},
		{
			name: "UpdateWithID",
			canvas: &storage.Canvas{
				ID: "test-canvas",
			},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			created, err := _testStorage.CreateCanvas(context.Background(), *test.canvas)
			if err != nil {
				if err != test.err {
					t.Fatalf("expected error %v got %v", test.err, err)
				}
			}
			if diff := cmp.Diff(test.canvas, created); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
