package test

import (
	ts "tdd/test/stocks"
	"testing"
)

/*
	功能清单
	1、 5 usd * 2 = 10 usd
	2、 10 EUR * 2 = 20 EUR
	3、 4002 KRW / 4 = 1000.5 KRW
	4、 5 USD + 10 EUR = 17 USD
	5、 将测试代码与产品代码分开
	6、 删除重复的测试
	7、 1 USD + 110 KRW = 2200 KRW
	8、 从与Money的乘法相关的那些测试方法中移除重复代码
*/

func TestAddition(t *testing.T) {
	var portfolio ts.Portfolio
	var portfolioInDollars ts.Money
	fiveDollars := ts.NewMoney(5, "USD")
	tenDollars := ts.NewMoney(10, "USD")
	fifteenDollars := ts.NewMoney(15, "USD")
	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars = portfolio.Evaluate(fiveDollars.Currency())
	assertEqual(t, fifteenDollars, portfolioInDollars)
}

func assertEqual(t *testing.T, expected, actual ts.Money) {
	if expected != actual {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}



func TestMultiplication(t *testing.T) {
	fiveEuros := ts.NewMoney(5, "EUR")
	actualResult := fiveEuros.Times(2)
	expectedResult := ts.NewMoney(10, "EUR")
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := ts.NewMoney(4002, "KRW")
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := ts.NewMoney(1000.5, "KRW")
	assertEqual(t, expectedMoneyAfterDivision, actualMoneyAfterDivision)
}
