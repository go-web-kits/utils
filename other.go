package utils

import (
	"reflect"
	"runtime"
	"strings"
)

func TypeNameOf(obj interface{}) string {
	name, ok := obj.(string)
	if !ok {
		name = reflect.Indirect(reflect.ValueOf(obj)).Type().String()
	}
	if i := strings.LastIndex(name, "."); i > -1 {
		name = name[(i + 1):]
	}
	return name
}

func GetFuncName(f interface{}) string {
	if s, ok := f.(string); ok {
		return s
	}
	path := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), "/")
	name := path[len(path)-1]
	return name
}
