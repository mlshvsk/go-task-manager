package models

import "time"

type Column struct {
	Id        int64     `json:"id" db:"notFillable"`
	Name      string    `json:"name,omitempty" validate:"required,max=255"`
	ProjectId int64     `json:"project_id" db:"notFillable"`
	Position  int64     `json:"position"`
	CreatedAt time.Time `json:"created_at" db:"notFillable"`
}
