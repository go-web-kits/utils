package logx

import (
	"fmt"
	"os"
	"time"

	"github.com/go-web-kits/utils/project"
)

func Log(flag, log string, lambdaOrDuration ...interface{}) {
	tip := getDuration(lambdaOrDuration...) + Blod(Magenta(flag))
	if project.OnDev() || os.Getenv("ENV") == "UAT" {
		args := []interface{}{"\n", Yello("[" + time.Now().Format("2006-01-02 15:04:05") + "]"), tip, log}
		fmt.Println(args...)
	}
}

func LogBy(client interface{ Println(args ...interface{}) }, flag, log string, lambdaOrDuration ...interface{}) {
	tip := getDuration(lambdaOrDuration...) + Blod(Magenta(flag))
	args := []interface{}{"\n", Yello("[" + time.Now().Format("2006-01-02 15:04:05") + "]"), tip, log}
	client.Println(args...)
}

// ======

func getDuration(lambdaOrDuration ...interface{}) string {
	if len(lambdaOrDuration) == 0 {
		return ""
	}

	var duration time.Duration
	switch x := lambdaOrDuration[0].(type) {
	case time.Duration:
		duration = x
	case func():
		start := time.Now()
		x()
		duration = time.Since(start)
	}

	return Blod(SkyBlue(float64(duration.Nanoseconds()/1e4)/100.0, "[%.2fms]")) + " "
}
