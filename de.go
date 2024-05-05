package main

func getDE(os string) string {
	switch os {
	case Darwin:
		return "Aqua"
	default:
		return ""
	}
}
