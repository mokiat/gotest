package match_test

import (
	"testing"

	. "github.com/mokiat/gotest"
	. "github.com/mokiat/gotest/match"
)

func TestEq(t *testing.T) {
	c := NewContext(t)

	matcher := Eq(5)
	Assert(c, matcher.Match(3).Error(), Eq("3 != 5"))
	Assert(c, matcher.Match(5), Eq[error](nil))
	Assert(c, matcher.Explain(), Eq("equal 5"))
}

func ExampleEq() {
	c := NewMockContext() // Would normally be NewContext(t)
	Assert(c, 10, Eq(10))
	Assert(c, "hello", Eq("hello"))
}
