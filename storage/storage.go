package storage

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Canvas struct {
	ID         string     `db:"id"`
	Operations Operations `db:"operations"`
}

type Operations struct {
	Rectangles []Rectangle `json:"rectangles"`
	Fills      []Fill      `json:"fills"`
}

type Rectangle struct {
	StartX  int    `json:"start_x"`
	StartY  int    `json:"start_y"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Outline string `json:"outline"`
	Fill    string `json:"fill"`
}

type Fill struct {
	StartX    int    `json:"start_x"`
	StartY    int    `json:"start_y"`
	Character string `json:"character"`
}

func (o Operations) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *Operations) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, o)
}
