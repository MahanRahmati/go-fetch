package main

import (
	"strings"
)

func getGPU(os string) string {
	switch os {
	case Darwin:
		out, errout, err := shellout("system_profiler SPDisplaysDataType")
		if err != nil {
			return ""
		}
		if errout != "" {
			return ""
		}
		if strings.Contains(out, "Chipset Model:") {
			start := strings.Index(out, "Chipset Model:") + 15
			end := strings.Index(out[start:], "\n")
			if end == -1 {
				end = len(out)
			}
			end += start
			gpu := out[start:end]
			return gpu
		}
		return ""
	default:
		return ""
	}
}
