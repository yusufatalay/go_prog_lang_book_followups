import "strings"

// funcrtions same as basename1 but this version uses string.LastIndex feature

func basename(s string) string {

	lastSlash := strings.LastIndex(s, "/") // returns -1 if there is no "/"
	s = s[lastSlash+1:]

	if lastDot := strings.LastIndex(s, "."); lastDot >= 0 {
		s = s[:lastDot]
	}

	return s
}
