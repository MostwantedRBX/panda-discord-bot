package crypt

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

//	TODO: make the GetCoins function data timestamped so the bot doesn't fetch it every time !coins is called.

var Coins []Coin

//	structs for formatting incoming coin json

type Coin struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Rank   int    `json:"rank"`
}

type USDQuote struct {
	Price    float32 `json:"price"`
	Last30   float32 `json:"percent_change_30m"`
	LastDay  float32 `json:"percent_change_24h"`
	LastWeek float32 `json:"percent_change_7d"`
}

type Quote struct {
	USD USDQuote `json:"USD"`
}

type CoinData struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	LastUpdate string `json:"last_updated"`
	Quotes     Quote  `json:"quotes"`
	Rank       int    `json:"rank"`
}

//	url to the crypto database api
var Url string = "https://api.coinpaprika.com/v1/"

func GetCoins() []Coin {
	//	fetch all the coins and their ids to compare to the argument in the command

	res, err := http.Get(Url + "coins")
	if err != nil {
		log.Logger.Err(err).Msg("Could not get coins from api... Is their server down?")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Logger.Err(err).Msg("Could not read coin data")
	}

	err = json.Unmarshal(body, &Coins)
	if err != nil {
		log.Logger.Err(err).Msg("Could not unmarshal json data")
	}

	return Coins
}

func FetchCoinData(c string) (CoinData, error) {
	//	get coin data for c(requested coin)
	var (
		coinStats CoinData
		coin      Coin
	)

	coins := GetCoins()

	for i, v := range coins {
		if strings.ToLower(v.Name) == c {
			coin = v
			break
		} else if i == len(coins)-1 {
			return CoinData{}, errors.New("the coin is not in the api")
		}
	}

	res, err := http.Get(Url + "tickers/" + coin.ID)
	if err != nil {
		return coinStats, errors.New("could not fetch coin ticker from api")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return coinStats, errors.New("could not read coin data")
	}

	err = json.Unmarshal(body, &coinStats)
	if err != nil {
		return coinStats, errors.New("could not unmarshal coin data json")
	}

	return coinStats, nil
}
