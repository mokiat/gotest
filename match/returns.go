package match

import "fmt"

func Returns[T any, F func() T](expectation Matcher[T]) Matcher[F] {
	match := func(fnActual F) error {
		actual := fnActual()
		return expectation.Match(actual)
	}
	explanation := fmt.Sprintf("returns %s", expectation.Explain())
	return NewMatcher[F](
		WithMatchFunc[F](match),
		WithExplanation[F](explanation),
	)
}
