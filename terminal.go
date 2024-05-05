package main

func getTerminal(os string) string {
	switch os {
	case Darwin:
		out, errout, err := shellout("echo \"$TERM_PROGRAM\"")
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
