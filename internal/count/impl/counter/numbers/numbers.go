package numbers

import "unicode"

type Numbers struct{}

func (n *Numbers) Exec(text string) int {
	count := 0
	for _, char := range text {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}
