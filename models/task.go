package models

import "time"

type Task struct {
	Id          int64     `json:"id" db:"notFillable"`
	Name        string    `json:"name" validate:"required,min=1,max=500"`
	Description string    `json:"description" validate:"max=5000"`
	ColumnId    int64     `json:"column_id" db:"notFillable"`
	Position    int64     `json:"position"`
	CreatedAt   time.Time `json:"created_at" db:"notFillable"`
}
