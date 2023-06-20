package example_test

import (
	"context"
	"math"
	"testing"

	. "github.com/mokiat/gotest/match"

	"github.com/mokiat/gotest/internal/example"
	"github.com/mokiat/gotest/internal/example/examplemock"
)

func TestConvertImplicit(t *testing.T) {
	exchangeClient := examplemock.NewClient()
	exchangeClient.WHEN().
		ExchangeRate(Any[context.Context](), Eq("BGN"), Eq("EUR")).
		Return(0.5, nil).
		Times(1)

	exchangeClient.WHEN().
		ExchangeRate(Any[context.Context](), Eq("USD"), Eq("EUR")).
		Return(0.9, nil).
		Times(1)

	converter := example.NewConverter(exchangeClient)

	amount, err := converter.Convert([]example.Source{
		{
			Currency: "BGN",
			Amount:   100.0,
		},
		{
			Currency: "USD",
			Amount:   150.0,
		},
	}, "EUR")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedAmount := 50.0 + 135.0
	if math.Abs(amount-expectedAmount) > 0.00001 {
		t.Errorf("wrong amount produced: %f", amount)
	}
}
