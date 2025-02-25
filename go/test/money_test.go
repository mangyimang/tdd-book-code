package test

import (
	ts "tdd/test/stocks"
	"testing"
	"reflect"
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
	9、 根据换算所涉及的币种确定汇率
	10、 改善错误处理机制，以应对相关汇率确实的情况
	11、 改进货币换算机制的实现方式
	12、 让汇率能够修改
*/

var bank ts.Bank

func init() {
	bank = ts.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	bank.AddExchangeRate("USD", "KRW", 1100)
}

func initExchangeRates() {
	bank = ts.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	bank.AddExchangeRate("USD", "KRW", 1100)
}


func TestConversionWithMissingExchangeRate(t *testing.T) {
	tenEuros := ts.NewMoney(10, "EUR")
	actualConvertedMoney, err := bank.Convert(tenEuros, "Kalganid")
	if actualConvertedMoney != nil {
		t.Errorf("Expected money to be nil, got %+v", actualConvertedMoney)
	}
	assertEqual(t,  "EUR->Kalganid", err.Error())
}

func TestConversion(t *testing.T) {
	tenEuros := ts.NewMoney(10, "EUR")
	actualResult, err := bank.Convert(tenEuros, "USD")
	assertNil(t, err)
	assertEqual(t, ts.NewMoney(12, "USD"), *actualResult)
	bank.AddExchangeRate("EUR", "USD", 1.3)
	actualResult, err = bank.Convert(tenEuros, "USD")
	assertNil(t, err)
	assertEqual(t, ts.NewMoney(13, "USD"), *actualResult)
}

func assertNil(t *testing.T, actual interface{}) {
	if actual != nil && !reflect.ValueOf(actual).IsNil() {
		t.Errorf("Expected nil, got [%+v]", actual)
	}
}

func TestAdditionWithMultipleMissingExchangeRates(t *testing.T) {
	var portfolio ts.Portfolio

	portfolio = portfolio.Add(ts.NewMoney(1, "USD"))
	portfolio = portfolio.Add(ts.NewMoney(1, "EUR"))
	portfolio = portfolio.Add(ts.NewMoney(1, "KRW"))

	expectedError := "Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid,]"
	value, actualError := portfolio.Evaluate(bank, "Kalganid")

	assertNil(t, value)
	assertEqual(t, expectedError, actualError.Error())
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
	var portfolio ts.Portfolio

	portfolio = portfolio.Add(ts.NewMoney(1, "USD"))
	portfolio = portfolio.Add(ts.NewMoney(1100, "KRW"))

	actualResult, err := portfolio.Evaluate(bank, "KRW")
	expectedResult := ts.NewMoney(2200, "KRW")
	assertNil(t, err)
	assertEqual(t, expectedResult, *actualResult)
}


func TestAdditionOfDollarsAndEuros(t *testing.T) {
	initExchangeRates()
	var portfolio ts.Portfolio
	fiveDollars := ts.NewMoney(5, "USD")
	tenEuros := ts.NewMoney(10, "EUR")
	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)
	actualResult, err := portfolio.Evaluate(bank, "USD")
	expectedResult := ts.NewMoney(17, "USD")
	assertNil(t, err)
	assertEqual(t, expectedResult, *actualResult)
}

func TestAddition(t *testing.T) {
	var portfolio ts.Portfolio

	fiveDollars := ts.NewMoney(5, "USD")
	tenDollars := ts.NewMoney(10, "USD")
	fifteenDollars := ts.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars, err := portfolio.Evaluate(bank, fiveDollars.Currency())

	assertNil(t, err)
	assertEqual(t, fifteenDollars, *portfolioInDollars)
}

func assertEqual(t *testing.T, expected, actual interface{}) {
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