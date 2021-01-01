package models

import "time"

type Column struct {
	Id        int       `json:"id" db:"notFillable"`
	Name      string    `json:"name,omitempty"`
	ProjectId int       `json:"project_id" db:"notFillable"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"created_at" db:"notFillable"`
}
