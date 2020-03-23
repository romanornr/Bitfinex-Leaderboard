package bitfinex

import "testing"

func TestGetTicker(t *testing.T) {
	_, err := GetBitcoinPrice()
	if err != nil {
		t.Fail()
	}
}
