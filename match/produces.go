package match

import "fmt"

func Produces[T any, C <-chan T](expectation Matcher[T]) Matcher[C] {
	var (
		receivedValue T
		didReceive    bool
	)
	match := func(ch C) error {
		if !didReceive {
			select {
			case value := <-ch:
				receivedValue = value
				didReceive = true
			default:
			}
		}
		if !didReceive {
			return fmt.Errorf("nothing produced")
		}
		return expectation.Match(receivedValue)
	}
	explanation := fmt.Sprintf("produces %s", expectation.Explain())
	return NewMatcher[C](
		WithMatchFunc[C](match),
		WithExplanation[C](explanation),
	)
}
