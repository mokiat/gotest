package gotest

import (
	"github.com/mokiat/gotest/format"
	"github.com/mokiat/gotest/match"
)

func Assert[T any](c C, value T, expectation match.Matcher[T]) {
	if err := expectation.Match(value); err != nil {
		c.Fail("Expectation not satisfied:\n%s", format.Indent(err.Error(), 1))
	}
}
