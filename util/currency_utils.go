package util

import "strings"
import "fmt"
import "strconv"

// FromCurrencied converts a currencied amount to component currency sign and amount
func FromCurrencied(currenciedAmount string) (string, float64, error) {
	amount := strings.Split(strings.TrimSpace(currenciedAmount), " ")
	if len(amount) != 2 {
		return "", 0, fmt.Errorf("currencied amount in wrong format %s. Should be KES 200", currenciedAmount)
	}

	currency, err := strconv.ParseFloat(amount[1], 64)
	if err != nil {
		return "", 0, fmt.Errorf("could not get currency amount: %v", err)
	}

	currencySign := strings.ToUpper(amount[0])

	return currencySign, currency, nil
}
