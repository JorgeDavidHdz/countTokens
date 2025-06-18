package count

type TokenCounter interface {
	Exec(text string) int
}
