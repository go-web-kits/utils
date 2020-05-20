package structx

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/imdario/mergo"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// func ToMap(obj interface{}) (m map[string]interface{}) {
// 	return m // TODO
// }

func ToJsonizeMap(obj interface{}) (m map[string]interface{}) {
	if o, ok := obj.(map[string]interface{}); ok {
		return o
	}
	j, _ := json.Marshal(obj)
	json.Unmarshal(j, &m)
	return m
}

func GetFieldOf(obj interface{}, fieldName string) reflect.StructField {
	field, exists := ReflectValueOfStruct(obj).Type().FieldByName(fieldName)
	if exists {
		return field
	} else {
		panic("GetFieldOf: no such field")
	}
}

func GetFieldValueOf(obj interface{}, fieldName string) interface{} {
	fieldChain := strings.Split(fieldName, ".")
	var field = ReflectValueOfStruct(obj)
	for _, name := range fieldChain {
		field = field.FieldByName(name)
	}

	if field.IsValid() {
		return field.Interface()
	} else {
		return nil
	}
}

func SetFieldValueOfTagName(obj interface{}, tag, tagName string, setVal interface{}) {
	val := reflect.Indirect(reflect.ValueOf(obj))
	tp := val.Type()
	for i := 0; i < tp.NumField(); i++ {
		tags := tp.Field(i).Tag.Get(tag)
		if tags == "" {
			continue
		}
		tagList := strings.Split(tags, ",")
		for _, t := range tagList {
			if t == tagName {
				val.Field(i).Set(reflect.ValueOf(setVal))
				return
			}
		}
	}

	return
}

// ToFieldName2TagValueMap(user, "db") => map[string]string{"Name": "name"}
func ToFieldName2TagValueMap(obj interface{}, tagName string) map[string]string {
	m := map[string]string{}
	val := ReflectValueOfStruct(obj)
	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Type().Field(i)
		m[field.Name] = field.Tag.Get(tagName)
	}
	return m
}

// ToFieldName2FieldTypeMap(user) => map[string]string{"Name": "string"}
func ToFieldName2FieldTypeMap(obj interface{}) map[string]string {
	m := map[string]string{}
	val := ReflectValueOfStruct(obj)
	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Type().Field(i)
		m[field.Name] = field.Type.String()
	}
	return m
}

// ToTagValue2FieldValueMap(user, "db") => map[string]string{"name": "Jack"}
func ToTagValue2FieldValueMap(obj interface{}, tagName string) map[string]interface{} {
	m := map[string]interface{}{}
	val := ReflectValueOfStruct(obj)
	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Type().Field(i)
		if field.Anonymous {
			fieldM := ToTagValue2FieldValueMap(val.Field(i).Interface(), tagName)
			for k, v := range fieldM {
				m[k] = v
			}
			continue
		}

		tag := strings.Split(field.Tag.Get(tagName), ",")
		key := tag[0]
		if key == "-" {
			continue
		}

		if len(tag) > 1 && tag[1] == "omitempty" {
			switch val.Field(i).Kind() {
			case reflect.Slice, reflect.Map:
				if val.Field(i).Len() == 0 {
					continue
				}
			default:
				if val.Field(i).Interface() == reflect.Zero(val.Field(i).Type()).Interface() {
					continue
				}
			}
		}
		if key == "" {
			key = field.Name
		}
		m[key] = val.Field(i).Interface() // if CanInterface()
	}
	return m
}

func Merge(objs ...interface{}) interface{} {
	if len(objs) == 0 {
		return false
	}

	dest := objs[0]
	for _, obj := range objs[1:] {
		// TODO error
		mergo.Merge(&dest, obj)
	}
	return dest
}

func ReflectValueOfStruct(obj interface{}) reflect.Value {
	o := reflect.Indirect(reflect.ValueOf(obj))

	switch o.Kind() {
	case reflect.Struct:
		return o
	default:
		// panic("")
		return reflect.ValueOf(struct{}{})
	}
}

func ToJsonb(obj interface{}) (dst postgres.Jsonb) {
	bs, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	err = dst.Scan(bs)
	if err != nil {
		panic(err)
	}

	return dst
}
