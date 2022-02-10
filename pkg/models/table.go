package models

import (
	"time"
)

type Table struct {
	Name      string         `json:"name"`
	TType     string         `json:"type"`
	Rows      int64          `json:"rows"`
	Schema    []*TableSchema `json:"schema"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type TableSchema struct {
	Field       string      `json:"field"`
	Type        string      `json:"type"`
	Null        string      `json:"null"`
	Key         string      `json:"key"`
	DefaultData interface{} `json:"default"`
	Extra       string      `json:"extra"`
}
