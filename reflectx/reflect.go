package reflectx

import "reflect"

func ValueOf(value interface{}, tp ...reflect.Type) reflect.Value {
	if e, ok := value.(error); ok {
		return reflect.ValueOf(&e).Elem()
	} else if value == nil {
		if len(tp) == 0 {
			panic("tp not given")
		}
		t := tp[0]
		return reflect.Indirect(reflect.Zero(t))
	} else {
		return reflect.ValueOf(value)
	}
}

func New(obj interface{}) interface{} {
	return reflect.New(reflect.Indirect(reflect.ValueOf(obj)).Type()).Interface()
}

func IsZero(obj interface{}) bool {
	return reflect.DeepEqual(reflect.Zero(reflect.TypeOf(obj)).Interface(), obj)
}
