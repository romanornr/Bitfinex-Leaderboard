package bitfinex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Ticker []float64

// return last price
func GetBitcoinPrice() (float64, error) {
	url := fmt.Sprintf("https://api-pub.bitfinex.com/v2/ticker/tBTCUSD")
	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	ticker := Ticker{}
	err = json.Unmarshal(body, &ticker)
	if err != nil {
		log.Fatal(err)
	}

	return ticker[6], err
}
