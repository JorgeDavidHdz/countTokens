package letters

import "unicode"

type Letters struct{}

func (l *Letters) Exec(text string) int {
	count := 0
	for _, char := range text {
		if unicode.IsLetter(char) {
			count++
		}
	}
	return count
}
