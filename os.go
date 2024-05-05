package main

import (
	"runtime"
	"strings"
)

const (
	Darwin string = "Darwin"
)

func getOS() string {
	os := runtime.GOOS
	if strings.Contains(os, "darwin") {
		return Darwin
	} else {
		return ""
	}
}
