package test

import (
	"testing"
)

/* 
	功能清单
	1、 5 usd * 2 = 10 usd
	2、 10 EUR * 2 = 20 EUR
	3、 4002 KRW / 4 = 1000.5 KRW
	4、 5 USD + 10 EUR = 17 USD
	5、 1 USD + 110 KRW = 2200 KRW
	6、 从与Money的乘法相关的那些测试方法中移除重复代码
*/

func TestAddition(t *testing.T) {
	var portfolio Portfolio
	var portfolioInDollars Money
	fiveDollars := Money{amount: 5, currency: "USD"}
	tenDollars := Money{amount: 10, currency: "USD"}
	fifteenDollars := Money{amount: 15, currency: "USD"}
	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars = portfolio.Evaluate(fiveDollars.currency)
	assertEqual(t, fifteenDollars, portfolioInDollars)
}

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

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	return append(p, money)
}

func (p Portfolio) Evaluate(currency string) Money {
	total := 0.0
	for _, money := range p {
		total = total + money.amount
	}
	return Money{amount: total, currency: currency}
}
