package handel

import (
	"strings"
)

// generateAscii converts the input text into ASCII art using the provided banner data.
// It handles multiple lines and prints the corresponding ASCII representation.
func GenerateAscii(input string, data []byte) string {
	datastr := strings.ReplaceAll(string(data), "\r", "")
	lines := strings.Split(datastr, "\n")
	var result []string 
	textlines := strings.Split(input, "\n")
	for _, text := range textlines {
		 if text == "" {
			result = append(result, "\n")
			continue
		}
		for row := 1; row < 9; row++ {
			for _, ch := range text {
				index := int(ch-32)*9 + row
				if index < len(lines) {
					result = append(result, lines[index])
				}
			}
				result = append(result, "\n")
		}
	}
	res := strings.Join(result, "")
	return res
}
