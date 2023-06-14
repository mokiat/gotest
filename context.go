package gotest

import (
	"fmt"
	"testing"
)

// C represents a testing context.
type C interface {
	Fail(format string, args ...any)
}

func NewContext(t *testing.T) *NativeC {
	return &NativeC{t: t}
}

type NativeC struct {
	t *testing.T
}

func (c *NativeC) Fail(format string, args ...any) {
	c.t.Errorf(format, args...)
}

func NewMockContext() *MockC {
	return &MockC{}
}

type MockC struct {
	failure string
}

func (c *MockC) Fail(format string, args ...any) {
	c.failure = fmt.Sprintf(format, args...)
}

func (c *MockC) DidFail() bool {
	return c.failure != ""
}

func (c *MockC) FailureMessage() string {
	return c.failure
}
