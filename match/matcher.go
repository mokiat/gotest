package match

import "fmt"

type Matcher[T any] interface {
	Explain() string
	Match(actual T) error
}

type MatcherOption[T any] func(m *matcher[T])

func WithMatchFunc[T any](fn func(actual T) error) MatcherOption[T] {
	return func(m *matcher[T]) {
		m.matchFunc = fn
	}
}

func WithExplanation[T any](explanation string) MatcherOption[T] {
	return func(m *matcher[T]) {
		m.explanation = explanation
	}
}

func NewMatcher[T any](opts ...MatcherOption[T]) Matcher[T] {
	result := &matcher[T]{
		matchFunc: func(actual T) error {
			return fmt.Errorf("match not implemented")
		},
		explanation: "Some expectation (not documented).",
	}
	for _, opt := range opts {
		opt(result)
	}
	return result
}

type matcher[T any] struct {
	matchFunc   func(actual T) error
	explanation string
}

func (m *matcher[T]) Match(actual T) error {
	return m.matchFunc(actual)
}

func (m *matcher[T]) Explain() string {
	return m.explanation
}
