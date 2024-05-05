package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func getUpTime(os string) string {
	switch os {
	case Darwin:
		out, errout, err := shellout("sysctl -n kern.boottime")
		if err != nil {
			return ""
		}
		if errout != "" {
			return ""
		}
		if strings.Contains(out, "sec =") {
			start := strings.Index(out, "sec =") + 6
			end := strings.Index(out, ",")
			if end == -1 {
				return ""
			}
			bootTime := out[start:end]
			boot, err := strconv.ParseInt(bootTime, 10, 64)
			if err != nil {
				return ""
			}
			return uptime(boot)
		}
		return ""
	default:
		return ""
	}
}

func uptime(boot int64) string {
	now := time.Now().Unix()
	uptime := now - boot
	d, h, m := duration(uptime)
	str := ""
	if d != 0 {
		str = str + fmt.Sprint(d) + " days"
		if h != 0 || m != 0 {
			str = str + ", "
		}
	}
	if h != 0 {
		str = str + fmt.Sprint(h) + " hours"
		if m != 0 {
			str = str + ", "
		}
	}
	if m != 0 {
		str = str + fmt.Sprint(m) + " mins"
	}
	return str
}

func duration(time int64) (days int64, hours int64, minutes int64) {
	if time < 60 {
		return 0, 0, 0
	}
	m := time / 60
	if m < 60 {
		return 0, 0, m
	}
	h := m / 60
	if h < 24 {
		return 0, h, m % 60
	}
	d := h / 24
	return d, h % 24, m % 60
}
