package models

import "time"

type Comment struct {
	Id 			int64		`json:"-" db:"notFillable"`
	Data 		string		`json:"data" validate:"required,max=5000"`
	TaskId 		int64		`json:"task_id" db:"notFillable"`
	CreatedAt 	time.Time	`json:"-" db:"notFillable"`
}
