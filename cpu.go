package main

func getCPU(os string) string {
	switch os {
	case Darwin:
		out, errout, err := shellout("sysctl -n machdep.cpu.brand_string")
		if err != nil {
			return ""
		}
		if errout != "" {
			return ""
		}
		return out
	default:
		return ""
	}
}
