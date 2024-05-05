package main

import (
	"fmt"
	"strings"
)

func main() {
	os := getOS()
	arch := getArch()
	model := getModel(os)
	kernel := getKernel(os)
	uptime := getUpTime(os)
	de := getDE(os)
	terminal := getTerminal(os)
	cpu := getCPU(os)
	gpu := getGPU(os)
	info := Info{
		os:       os,
		arch:     arch,
		model:    model,
		kernel:   kernel,
		uptime:   uptime,
		de:       de,
		terminal: terminal,
		cpu:      cpu,
		gpu:      gpu,
	}
	printInformation(info)
}

func printInformation(info Info) {
	println("OS", info.os)
	println("Architecture", info.arch)
	println("Host", info.model)
	println("Kernel", info.kernel)
	println("Uptime", info.uptime)
	println("DE", info.de)
	println("Terminal", info.terminal)
	println("CPU", info.cpu)
	println("GPU", info.gpu)
}

func println(title string, data string) {
	if strings.Trim(data, " ") == "" {
		return
	}
	d := strings.TrimSuffix(data, "\n")
	fmt.Println(title + ": " + d)
}
