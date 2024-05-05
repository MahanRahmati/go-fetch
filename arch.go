package main

import (
	"runtime"
)

func getArch() string {
	arch := runtime.GOARCH
	return arch
}
