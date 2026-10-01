package main

func isValid(s string) bool {
	return false
}

func getClosing(s string) string {
	if s == "(" {
		return ")"
	} else if s == "{" {
		return "}"
	} else if s == "[" {
		return "]"
	}
	return ""
}
