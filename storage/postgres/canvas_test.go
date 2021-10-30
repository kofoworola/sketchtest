package postgres

import (
	"context"
	"database/sql"
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
			created, err := _testStorage.CreateOrUpdateCanvas(context.Background(), *test.canvas)
			if err != test.err {
				t.Fatalf("expected error %v got %v", test.err, err)
			}
			if diff := cmp.Diff(test.canvas, created); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestFetchCanvas(t *testing.T) {
	canv := storage.Canvas{
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
	}
	if _, err := _testStorage.CreateOrUpdateCanvas(context.Background(), canv); err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		Name     string
		id       string
		err      error
		expected *storage.Canvas
	}{
		{
			Name:     "SuccessfulFetch",
			id:       "test-canvas",
			err:      nil,
			expected: &canv,
		},
		{
			Name:     "FetchNotFound",
			id:       "dummy-canvas",
			err:      sql.ErrNoRows,
			expected: nil,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			gotten, err := _testStorage.GetCanvasById(context.Background(), test.id)
			if err != test.err {
				t.Fatalf("expected error %v gotten %v", test.err, err)
			}

			if diff := cmp.Diff(test.expected, gotten); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
