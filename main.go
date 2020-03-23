package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type BitfinexLeaderboard [][]interface{}

type LeaderboardData struct {
	Name             string
	Rank             float64
	UnrealizedProfit float64
	RealizedProfit   float64
	Twitter          string
	Date             time.Time
}

func main() {

	url := fmt.Sprintf("https://api-pub.bitfinex.com/v2/rankings/plu_diff:1w:tGLOBAL:USD/hist")
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

	BitfinexLeaderboard := BitfinexLeaderboard{}
	err = json.Unmarshal(body, &BitfinexLeaderboard)
	if err != nil {
		log.Fatal(err)
	}

	leaderboard := LeaderboardData{
		Name:             fmt.Sprintf("%s", BitfinexLeaderboard[0][2]),
		Rank:             BitfinexLeaderboard[0][3].(float64),
		UnrealizedProfit: BitfinexLeaderboard[0][6].(float64),
		RealizedProfit:   0,
		Twitter:          fmt.Sprintf("%s", BitfinexLeaderboard[0][9]),
		Date:             time.Now(),
	}

	fmt.Println(leaderboard)
}
