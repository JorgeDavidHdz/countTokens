package sentences

import (
	"regexp"
	"strings"
)

type Sentences struct{}

func (s *Sentences) Exec(text string) int {
	sentenceRegex := regexp.MustCompile(`[.!?]+`)
	sentences := sentenceRegex.Split(text, -1)
	count := 0
	for _, sentence := range sentences {
		if strings.TrimSpace(sentence) != "" {
			count++
		}
	}
	return count
}
