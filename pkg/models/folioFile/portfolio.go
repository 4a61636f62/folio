package folioFile

import (
	"encoding/json"
	"folio/pkg/models"
	"os"
)

type PortfolioFile struct {
	filepath string
	portfolio *models.Portfolio
}

func Portfolio(filepath string) *PortfolioFile {
	return &PortfolioFile{
		filepath: filepath,
		portfolio: &models.Portfolio{},
	}
}

func (p *PortfolioFile) Holdings() (map[string]float32, error) {
	err := p.load()
	if err != nil {
		return nil, err
	}
	return p.portfolio.Holdings, nil
}

func (p *PortfolioFile) Add(coinId string, amount float32) error {
	err := p.load()
	if err != nil {
		return err
	}
	p.portfolio.Holdings[coinId] += amount
	return p.save()
}

func (p *PortfolioFile) Remove(coinId string) error {
	err := p.load()
	if err != nil {
		return err
	}
	delete(p.portfolio.Holdings, coinId)
	return p.save()
}


func (p *PortfolioFile) load() error {
	f, err := os.Open(p.filepath)
	if err != nil {
		p.portfolio.Holdings = make(map[string]float32)
		return nil
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	return dec.Decode(&p.portfolio.Holdings)
}

func (p *PortfolioFile) save() error {
	f, err := os.OpenFile(p.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	return enc.Encode(&p.portfolio.Holdings)
}


