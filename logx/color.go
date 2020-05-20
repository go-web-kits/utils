// https://toutiao.io/posts/2889gp/preview
package logx

import "fmt"

func Blod(s string) string {
	return "\033[1m" + s + "\033[0m"
}

func Red(s interface{}, format ...string) string {
	if len(format) == 0 {
		format = append(format, "%s")
	}
	return fmt.Sprintf("\x1b[91m"+format[0]+"\x1b[0m", s)
}

// 洋红
func Magenta(s interface{}, format ...string) string {
	if len(format) == 0 {
		format = append(format, "%s")
	}
	return fmt.Sprintf("\x1b[95m"+format[0]+"\x1b[0m", s)
}

func Yello(s interface{}, format ...string) string {
	if len(format) == 0 {
		format = append(format, "%s")
	}
	return fmt.Sprintf("\033[33m"+format[0]+"\033[0m", s)
}

func SkyBlue(s interface{}, format ...string) string {
	if len(format) == 0 {
		format = append(format, "%s")
	}
	return fmt.Sprintf("\033[36m"+format[0]+"\033[0m", s)
}
