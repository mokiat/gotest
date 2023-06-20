package example

import "context"

// TODO: Add a go:generate command to build the mock for this.

type Client interface {
	ExchangeRate(ctx context.Context, fromCurrency, toCurrency string) (amount float64, err error)
}
