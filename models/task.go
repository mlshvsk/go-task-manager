package models

import "time"

type Task struct {
	Id 				int			`json:"id" db:"notFillable"`
	Name 			string		`json:"name"`
	Description 	string		`json:"description"`
	ColumnId 		int			`json:"column_id" db:"notFillable"`
	Position  		int			`json:"position"`
	CreatedAt 		time.Time	`json:"created_at" db:"notFillable"`
}
