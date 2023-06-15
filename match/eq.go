package match

import "fmt"

func Eq[T comparable](expected T) Matcher[T] {
	return NewMatcher[T](
		WithMatchFunc[T](func(actual T) error {
			if actual != expected {
				return fmt.Errorf("%v != %v", actual, expected)
			}
			return nil
		}),
		WithExplanation[T](fmt.Sprintf("equal %v", expected)),
	)
}
