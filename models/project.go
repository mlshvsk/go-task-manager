package models

import "time"

type Project struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name,omitempty" validate:"required,max=500"`
	Description string    `json:"description,omitempty" validate:"max=1000"`
	CreatedAt   time.Time `json:"created_at"`
}
