package counting

import (
	"countTokens/internal/count"
	"countTokens/internal/count/counting/factory"
)

type ServiceCountToken interface {
	CountTokenProcess(text string, inputCategory []string, factory *factory.TokenFactory) (count.Token, error)
}
