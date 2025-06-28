package factory

import (
	"countTokens/internal/count"
	"countTokens/internal/count/impl/counter/letters"
	"countTokens/internal/count/impl/counter/numbers"
	"countTokens/internal/count/impl/counter/sentences"
	"countTokens/internal/count/impl/counter/symbols"
	"countTokens/internal/count/impl/counter/words"
	"fmt"
)

type TokenFactory struct{}

func (tf *TokenFactory) GetCounterFactory(category string) (count.TokenCounter, error) {
	switch category {
	case "letters":
		return &letters.Letters{}, nil
	case "symbols":
		return &symbols.Symbols{}, nil
	case "numbers":
		return &numbers.Numbers{}, nil
	case "words":
		return &words.Words{}, nil
	case "sentences":
		return &sentences.Sentences{}, nil
	default:
		return nil, fmt.Errorf("invalid category: %s", category)
	}
}
