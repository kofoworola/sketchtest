package postgres

import (
	"context"
	"fmt"

	"github.com/kofoworola/sketchtest/storage"
)

const insertCanvasSql = `
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

func (s *Storage) CreateCanvas(ctx context.Context, input storage.Canvas) (*storage.Canvas, error) {
	var canvas storage.Canvas

	stmt, err := s.db.PrepareNamedContext(ctx, insertCanvasSql)
	if err != nil {
		return nil, fmt.Errorf("error preparing statement: %w", err)
	}

	if err := stmt.Get(&canvas, input); err != nil {
		return nil, fmt.Errorf("error inserting into table: %w", err)
	}
	return &canvas, nil

}
