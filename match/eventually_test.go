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
	Assert(c, matcher.Match(3).Error(), Eq("still 3 != 5"))
	Assert(c, matcher.Match(5), Eq[error](nil))
	Assert(c, matcher.Explain(), Eq("eventually equal 5"))
}

func ExampleEventually_func() {
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
	}, Eventually(Returns(Eq(int64(5)))))
}

func ExampleEventually_chan() {
	c := NewMockContext() // Would normally be NewContext(t)

	ch := make(chan string)
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "hello"
	}()

	Assert(c, ch, Eventually(Produces(Eq("hello"))))
}
