package project

import (
	"os"
)

// func Root() string {
// 	// if OnTest() {
// 	// 	// out, _ := exec.Command("pwd").Output()
// 	// 	// pp.Println(string(out))
// 	// 	return os.Args[0] + "/"
// 	// }
//
// 	path, _ := os.Getwd()
// 	return regexp.MustCompile(`.*code\.xx\.com/[^/]*/[^/]*`).FindString(path) + "/"
// }

func OnTest() bool {
	return os.Getenv("TEST") == "true"
}

func OnDev() bool {
	return os.Getenv("ENV") == "DEV" && !OnTest()
}

func OnDevOrTest() bool {
	return OnTest() || OnDev()
}
