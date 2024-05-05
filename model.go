package main

func getModel(os string) string {
	switch os {
	case Darwin:
		out, errout, err := shellout("sysctl -n hw.model")
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
