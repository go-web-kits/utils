package slicex

import (
	"reflect"
	"strings"
)

func MapStrToI(s []string, lambda func(string) interface{}) []interface{} {
	result := []interface{}{}
	for _, item := range s {
		result = append(result, lambda(item))
	}
	return result
}

func MapStrToStr(s []string, lambda func(string) string) []string {
	result := []string{}
	for _, item := range s {
		result = append(result, lambda(item))
	}
	return result
}

func Append() {
	//
}

func GrepStr(strs []string, substr string) []string {
	result := []string{}
	for _, s := range strs {
		if strings.Contains(s, substr) {
			result = append(result, s)
		}
	}
	return result
}

func RemoveStr(strs []string, s string) []string {
	result := []string{}
	for _, item := range strs {
		if item != s {
			result = append(result, item)
		}
	}
	return result
}

func IncludeStr(strs []string, s string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}

func IncludeUint(nums []uint, num uint) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

func Reduce(maps []map[string]string, key string) []string {
	result := []string{}
	for _, m := range maps {
		result = append(result, m[key])
	}
	return result
}

func RejectIf(slice interface{}, lambda func(interface{}) bool) interface{} {
	switch s := slice.(type) {
	case []map[string]interface{}:
		result := []map[string]interface{}{}
		for _, item := range s {
			if !lambda(item) {
				result = append(result, item)
			}
		}
		return result
	case *[]map[string]interface{}:
		result := []map[string]interface{}{}
		for _, item := range *s {
			if !lambda(item) {
				result = append(result, item)
			}
		}
		return result
	default:
		panic("RejectIf")
	}
}

func Each(obj interface{}, lambda func(item interface{})) {
	refVal := reflect.Indirect(reflect.ValueOf(obj))
	if refVal.Kind() == reflect.Slice {
		for i := 0; i < refVal.Len(); i++ {
			lambda(refVal.Index(i).Interface())
		}
	} else {
		lambda(obj)
	}
}
