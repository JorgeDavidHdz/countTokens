package count

type TokenCounter interface {
	Exec(text string) int
}
type Token map[string]int
