package postgres

import (
	"context"
	"fmt"

	"github.com/kofoworola/sketchtest/storage"
)

const insertCanvasSQL = `
	INSERT INTO canvas (
		id, 
		operations
	) VALUES (
		:id, 
		:operations
	) ON CONFLICT ON CONSTRAINT canvas_pkey
	DO UPDATE SET operations = EXCLUDED.operations
	RETURNING *
`

const fetchCanvasSQL = `SELECT * FROM canvas WHERE id = $1`

func (s *Storage) CreateOrUpdateCanvas(ctx context.Context, input storage.Canvas) (*storage.Canvas, error) {
	var canvas storage.Canvas

	stmt, err := s.db.PrepareNamedContext(ctx, insertCanvasSQL)
	if err != nil {
		return nil, fmt.Errorf("error preparing statement: %w", err)
	}

	if err := stmt.Get(&canvas, input); err != nil {
		return nil, fmt.Errorf("error inserting into table: %w", err)
	}
	return &canvas, nil
}

func (s *Storage) GetCanvasById(ctx context.Context, id string) (*storage.Canvas, error) {
	var canvas storage.Canvas
	if err := s.db.Get(&canvas, fetchCanvasSQL, id); err != nil {
		return nil, err
	}
	return &canvas, nil
}
