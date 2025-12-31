package main

import (
	"strings"
)

// generateAscii converts the input text into ASCII art using the provided banner data.
// It handles multiple lines and prints the corresponding ASCII representation.
func generateAscii(input string, data []byte) []string {
	datastr := strings.ReplaceAll(string(data), "\r", "")

	lines := strings.Split(datastr, "\n")
	var result []string 
	onlynewline := true
	textlines := strings.Split(input, "\n")
	for i, text := range textlines {
		if text != "" {
			onlynewline = false
		} else if text == "" && !(onlynewline && i == 0) {
			result = append(result, "\n")
			continue
		}
		for row := 1; row < 9; row++ {
			// var line string
			for _, ch := range text {
				index := int(ch-32)*9 + row
				if index < len(lines) {
					result = append(result, lines[index])
					// line += lines[index]
				}
			}

			if !onlynewline {
				result = append(result, "\n")
			}
		}
	}

	return result
}
