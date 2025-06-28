package counting

import (
	"countTokens/internal/count"
	"countTokens/internal/count/counting/factory"
)

type Service struct{}

func (s *Service) CountTokenProcess(text string, inputCategory []string, factory *factory.TokenFactory) (count.Token, error) {
	result := make(count.Token)

	for _, category := range inputCategory {
		counter, err := factory.GetCounterFactory(category)
		if err != nil {
			return result, err
		}
		result[category] = counter.Exec(text)
	}

	return result, nil
}
