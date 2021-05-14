// basename removes directory  components and a .suffix.

func basename(s string) string {

	// Discard lasr '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// preserve everything before last '.'
	for i := len(s); i >= 0; i-- {
		if s[i] == "." {
			s = s[:i]
			break
		}
	}
	return s
}
