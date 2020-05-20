package mapx

func Merge(maps ...interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for _, m := range maps {
		switch _m := m.(type) {
		case map[string]interface{}:
			for k, v := range _m {
				result[k] = v
			}
		case map[string]string:
			for k, v := range _m {
				result[k] = v
			}
		default:
			panic("Merge")
		}
	}
	return result
}

func Keys(m map[string]interface{}) []string {
	var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values(m map[string]interface{}) []interface{} {
	var values []interface{}
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func Slice(m map[string]interface{}, keys []string) map[string]interface{} {
	result := map[string]interface{}{}
	for _, k := range keys {
		result[k] = m[k]
	}
	return result
}

func ValuesOf(m map[string]interface{}, keys []string) []interface{} {
	var values []interface{}
	for _, k := range keys {
		values = append(values, m[k])
	}
	return values
}

func Copy(i interface{}) interface{} {
	switch _i := i.(type) {
	case map[string]string:
		result := map[string]string{}
		for k, v := range _i {
			result[k] = v
		}
		return result
	default:
		return nil
	}
}

func Delete(m *map[string]interface{}, key string) interface{} {
	ret := (*m)[key]
	delete(*m, key)
	return ret
}

func Except(m map[string]interface{}, key ...string) map[string]interface{} {
	duplicateMap := map[string]interface{}{}
	for k, v := range m {
		duplicateMap[k] = v
	}
	for _, value := range key {
		delete(duplicateMap, value)
	}
	return duplicateMap
}

func Dig(m map[string]interface{}, keys ...string) interface{} {
	nxt := m[keys[0]]
	if len(keys[1:]) == 0 {
		return nxt
	}
	nxtVal, isMap := nxt.(map[string]interface{})
	if isMap {
		return Dig(nxtVal, keys[1:]...)
	} else {
		return nil
	}
}
