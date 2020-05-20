package stringx

func Slice(s string, from int, to int) string {
	if to < 0 {
		to = len(s) + to
	}
	if from < 0 {
		from = len(s) + from
	}
	if from > len(s) || to > len(s) || to < from {
		return s
	}

	return s[from:to]
}

func TheFirstNotEmpty(strs ...string) string {
	for _, s := range strs {
		if s != "" {
			return s
		}
	}
	return ""
}

func TheFirstNotNil(args ...interface{}) string {
	for _, arg := range args {
		if arg == nil {
			continue
		} else if s, ok := arg.(string); ok {
			return s
		}
	}
	return ""
}
