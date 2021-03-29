package folioFile

import (
	"encoding/json"
	"folio/pkg/coingecko"
	"folio/pkg/models"
	"os"
)

type CoinsFile struct {
	filepath string
	coins map[string]models.Coin
}

func Coins(filepath string) *CoinsFile {
	return &CoinsFile{
		filepath: filepath,
	}
}

func (c *CoinsFile) Update() error {
	coinList, err := coingecko.CoinsList()
	if err != nil {
		return err
	}
	c.coins = make(map[string]models.Coin)
	for _, coin := range coinList {
		c.coins[coin.Id] = coin
	}
	return c.save()
}

func (c *CoinsFile) All() (map[string]models.Coin, error) {
	err := c.load()
	if err != nil {
		return nil, err
	}
	return c.coins, nil
}

func (c *CoinsFile) load() error {
	f, err := os.Open(c.filepath)
	if err != nil {
		c.coins = make(map[string]models.Coin)
		return nil
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	return dec.Decode(&c.coins)
}

func (c *CoinsFile) save() error {
	f, err := os.OpenFile(c.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	return enc.Encode(c.coins)
}

