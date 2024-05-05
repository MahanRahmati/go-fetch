package main

func getKernel(os string) string {
	switch os {
	case Darwin:
		out, errout, err := shellout("uname -r")
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
