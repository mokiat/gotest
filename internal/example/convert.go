package example

type Source struct {
	Currency string
	Amount   float64
}

func NewConverter(exchangeClient Client) *Converter {
	return &Converter{
		exchangeClient: exchangeClient,
	}
}

type Converter struct {
	exchangeClient Client
}

func (c *Converter) Convert(sources []Source, targetCurrency string) (float64, error) {
	return 0.0, nil
}
