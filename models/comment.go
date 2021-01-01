package models

import "time"

type Comment struct {
	Id 			int			`json:"-" db:"notFillable"`
	Data 		string		`json:"data"`
	TaskId 		int			`json:"task_id" db:"notFillable"`
	CreatedAt 	time.Time	`json:"-" db:"notFillable"`
}
