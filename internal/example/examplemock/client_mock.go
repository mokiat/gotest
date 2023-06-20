package examplemock

import (
	"context"
	"fmt"

	"github.com/mokiat/gotest/internal/example"
	"github.com/mokiat/gotest/match"
	"github.com/mokiat/gotest/mock"
)

func NewClient() *Client {
	return &Client{
		registry: mock.NewRegistry(),
	}
}

type Client struct {
	registry *mock.Registry
}

var _ example.Client = (*Client)(nil)

func (c *Client) WHEN() *ClientWhen {
	return &ClientWhen{
		registry: c.registry,
	}
}

func (c *Client) ExchangeRate(ctx context.Context, fromCurrency, toCurrency string) (float64, error) {
	input := []any{ctx, fromCurrency, toCurrency}
	output := c.registry.Invoke("ExchangeRate", input)
	if len(output) == 0 {
		panic(fmt.Errorf("no precondition matches request parameters %#v", input))
	}
	return output[0].(float64), output[1].(error)
}

type ClientWhen struct {
	registry *mock.Registry
}

func (c *ClientWhen) ExchangeRate(ctx match.Matcher[context.Context], fromCurrency match.Matcher[string], toCurrency match.Matcher[string]) *ClientWhenExchangeRate {
	invocation := c.registry.Match("ExchangeRate", func(ctxActual context.Context, fromCurrencyActual string, toCurrencyActual string) bool {
		if err := ctx.Match(ctxActual); err != nil {
			return false
		}
		if err := fromCurrency.Match(fromCurrencyActual); err != nil {
			return false
		}
		if err := toCurrency.Match(toCurrencyActual); err != nil {
			return false
		}
		return true
	})
	return &ClientWhenExchangeRate{
		invocation: invocation,
	}
}

type ClientWhenExchangeRate struct {
	invocation *mock.Invocation
}

func (c *ClientWhenExchangeRate) Do(cb func(ctx context.Context, fromCurrency, toCurrency string) (amount float64, err error)) *ClientWhenExchangeRate {
	c.invocation.SetStub(cb)
	return c
}

func (c *ClientWhenExchangeRate) Return(amount float64, err error) *ClientWhenExchangeRate {
	output := []any{amount, err}
	c.invocation.SetReturnValues(output)
	return c
}

func (c *ClientWhenExchangeRate) Times(count int) *ClientWhenExchangeRate {
	c.invocation.SetCallCount(count)
	return c
}

func (c *ClientWhenExchangeRate) ForgetAboutIt() *ClientWhenExchangeRate {
	c.Times(0)
	return c
}
