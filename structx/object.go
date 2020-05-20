package structx

import (
	"strings"

	"github.com/go-web-kits/utils/mapx"
	"github.com/iancoleman/strcase"
)

type Struct struct {
	Raw interface{}
}

func From(obj interface{}) *Struct {
	return &Struct{Raw: obj}
}

func (s *Struct) ToMap(tagy ...string) mapx.Map {
	m := map[string]interface{}{}
	obj := s.Raw
	val := ReflectValueOfStruct(obj)
	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Type().Field(i)
		key := strcase.ToSnake(field.Name)

		if len(tagy) > 0 {
			tag := strings.Split(field.Tag.Get(tagy[0]), ",")
			key = tag[0]
			if key == "-" {
				continue
			}
		}

		m[key] = val.Field(i).Interface() // if CanInterface()
	}

	return m
}
