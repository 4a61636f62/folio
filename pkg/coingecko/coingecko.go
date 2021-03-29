package coingecko

import (
	"encoding/json"
	"folio/pkg/models"
	"net/http"
	"strings"
)

var base string = "https://api.coingecko.com/api/v3/"

func CoinsList() ([]models.Coin, error) {
	var ret []models.Coin
	resp, err := http.Get(base +"coins/list")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func SimplePrice(coinIDs, vsCurrencies []string) (map[string]map[string]float32, error) {
	ret := make(map[string]map[string]float32)
	resp, err := http.Get(base +
		"simple/price?ids=" + strings.Join(coinIDs, "%2C") +
		"&vs_currencies=" + strings.Join(vsCurrencies, "%2C"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}