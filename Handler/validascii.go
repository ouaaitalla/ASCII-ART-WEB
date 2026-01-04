package handel


// isValidASCII checks whether the string contains only printable ASCII characters or newlines.
func IsValidASCII(s string) bool {
	for _, r := range s {
		if r == '\n' {
			continue
		}
		if r >= 32 && r <= 126 {
			continue
		}
		return false
	}
	return true
}
