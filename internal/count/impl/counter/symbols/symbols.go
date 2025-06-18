package symbols

import "unicode"

type Symbols struct{}

func (s *Symbols) Exec(text string) int {
	count := 0
	for _, char := range text {
		if unicode.IsPunct(char) {
			count++
		}
	}
	return count
}
