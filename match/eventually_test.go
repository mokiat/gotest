package match_test

import (
	"sync/atomic"
	"testing"
	"time"

	. "github.com/mokiat/gotest"
	. "github.com/mokiat/gotest/match"
)

func TestEventually(t *testing.T) {
	c := NewContext(t)

	matcher := Eventually(Eq(5))
	Assert(c, matcher.Match(func() int { return 3 }).Error(), Eq("still 3 != 5"))
	Assert(c, matcher.Match(func() int { return 5 }), Eq[error](nil))
	Assert(c, matcher.Explain(), Eq("eventually equal 5"))
}

func ExampleEventually() {
	c := NewMockContext() // Would normally be NewContext(t)

	var value int64
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(200 * time.Millisecond)
			atomic.AddInt64(&value, 1)
		}
	}()

	Assert(c, func() int64 {
		return atomic.LoadInt64(&value)
	}, Eventually(Eq(int64(5))))
}
