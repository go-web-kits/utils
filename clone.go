package utils

import (
	"reflect"
	"time"
)

func Clone(src interface{}) interface{} {
	srcVal := reflect.Indirect(reflect.ValueOf(src))
	dstVal := reflect.New(srcVal.Type())
	reflect.Indirect(dstVal).Set(srcVal)
	return dstVal.Interface()
}

func Dup(obj interface{}) interface{} {
	newObj := Clone(obj)
	refVal := reflect.Indirect(reflect.ValueOf(newObj))
	switch refVal.Kind() {
	case reflect.Struct:
		id, createdAt, updatedAt :=
			refVal.FieldByName("ID"), refVal.FieldByName("CreatedAt"), refVal.FieldByName("UpdatedAt")
		if id.IsValid() {
			id.Set(reflect.ValueOf(uint(0)))
		}
		if createdAt.IsValid() {
			createdAt.Set(reflect.Zero(reflect.TypeOf(time.Now())))
		}
		if updatedAt.IsValid() {
			updatedAt.Set(reflect.Zero(reflect.TypeOf(time.Now())))
		}
	}

	return refVal.Interface()
}
