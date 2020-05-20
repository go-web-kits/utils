package mapx

import "reflect"

type Map map[string]interface{}

func (m Map) Compact() Map {
	for k, v := range m {
		if v == nil {
			delete(m, k)
			continue
		}

		if reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface()) {
			delete(m, k)
			continue
		}
	}
	return m
}

func (m Map) Except(keys ...string) Map {
	for _, k := range keys {
		delete(m, k)
	}
	return m
}

func (m Map) Merge(others ...map[string]interface{}) Map {
	slice := []interface{}{map[string]interface{}(m)}
	for _, v := range others {
		slice = append(slice, v)
	}
	return Merge(slice...)
}
