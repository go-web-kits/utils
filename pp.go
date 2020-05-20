package utils

import "github.com/k0kubun/pp"

func P(obj interface{}) interface{} {
	pp.Println(obj)
	return obj
}
