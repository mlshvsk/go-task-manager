package helpers

import (
	"github.com/mlshvsk/go-task-manager/models"
	"reflect"
)

func PrepareUpdatedModel(m models.Model) (fieldNames []string, values []interface{}) {
	t := reflect.TypeOf(m)
	for i := 0; i < t.NumField(); i++ {
		fieldNames = append(fieldNames, t.Field(i).Tag.Get("db"))
	}

	v := reflect.ValueOf(m)
	for i := 0; i < v.NumField(); i++ {
		values = append(values, v.Field(i))
	}

	return
}
