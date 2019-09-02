package util

import (
	"fmt"
	"os"
	"strings"
)

const DefaultErrorExitCode = 1

func CheckErr(err error) {
	checkErr(err, fatal)
}

func checkErr(err error, handleErr func(string, int)) {
	switch err.(type) {
	case nil:
		return
	default:
		handleErr(err.Error(), DefaultErrorExitCode)
	}
}

func fatal(msg string, code int) {
	if len(msg) > 0 {
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}

		fmt.Fprint(os.Stderr, msg)
	}
	os.Exit(code)
}
