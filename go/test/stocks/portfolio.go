package stocks

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	return append(p, money)
}

func (p Portfolio) Evaluate(currency string) Money {
	total := 0.0
	for _, money := range p {
		total = total + convert(money, currency)
	}
	return Money{amount: total, currency: currency}
}

func convert(money Money, currency string) float64 {
	exchangeRates := map[string]float64{
		"EUR->USD": 1.2,
		"USD->KRW": 1100,
	}

	if money.Currency() == currency {
		return money.amount
	}

	key := money.Currency() + "->" + currency
	rate := exchangeRates[key]
	return money.amount * rate
}
