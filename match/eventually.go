package match

import (
	"fmt"
	"time"
)

func Eventually[T any, F func() T](expectation Matcher[T]) Matcher[F] {
	const (
		interval = 100 * time.Millisecond
		timeout  = time.Second
	)

	match := func(fnActual F) error {
		deadline := time.Now().Add(timeout)
		// FIXME: This leaks memory
		for {
			actual := fnActual()
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

	return NewMatcher[F](
		WithMatchFunc[F](match),
		WithExplanation[F](explanation),
	)
}
