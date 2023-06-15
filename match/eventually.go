package match

import (
	"fmt"
	"time"
)

func Eventually[T any](expectation Matcher[T]) Matcher[T] {
	const (
		interval = 100 * time.Millisecond
		timeout  = time.Second
	)

	match := func(actual T) error {
		deadline := time.Now().Add(timeout)
		// FIXME: This leaks memory
		for {
			err := expectation.Match(actual)
			if err == nil {
				return nil
			}
			if time.Now().After(deadline) {
				return fmt.Errorf("still %s", err)
			}
			time.Sleep(interval)
		}
	}

	explanation := fmt.Sprintf("eventually %s", expectation.Explain())

	return NewMatcher[T](
		WithMatchFunc[T](match),
		WithExplanation[T](explanation),
	)
}
