package test

import (
	"testing"
)

func assertEqual(t *testing.T, expected, actual Money) {
	if expected != actual {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestMultiplicationInDollars(t *testing.T) {
	fiver := Money{
		amount: 5,
		currency: "USD",
	}
	actualResult := fiver.Times(2)
	expectedResult := Money{amount: 10, currency: "USD"}
	assertEqual(t, expectedResult, actualResult)
}

func TestMultiplicationInEuros(t *testing.T) {
	fiveEuros := Money{amount: 5, currency: "EUR"}
	actualResult := fiveEuros.Times(2)
	expectedResult := Money{amount: 10, currency: "EUR"}
	assertEqual(t, expectedResult, actualResult)
}


func TestDivision(t *testing.T) {
	originalMoney := Money{amount: 4002, currency: "KRW"}
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := Money{amount: 1000.5, currency: "KRW"}
	assertEqual(t, expectedMoneyAfterDivision, actualMoneyAfterDivision)
}
	


type Money struct {
	amount   float64
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * float64(multiplier), currency: m.currency}
}

func (m Money) Divide(divisor int) Money {
	return Money{amount: m.amount / float64(divisor), currency: m.currency}
}
