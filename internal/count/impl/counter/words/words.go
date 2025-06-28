package words

import "strings"

type Words struct{}

func (w *Words) Exec(text string) int {
	words := strings.Fields(text)
	return len(words)
}
