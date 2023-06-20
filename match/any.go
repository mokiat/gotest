package match

func Any[T comparable]() Matcher[T] {
	return NewMatcher[T](
		WithMatchFunc[T](func(actual T) error {
			return nil
		}),
		WithExplanation[T]("whichever"),
	)
}
