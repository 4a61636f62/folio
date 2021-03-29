package models

type Coin struct {
	Id string
	Name string
	Symbol string
}

type Portfolio struct {
	Holdings map[string]float32
}
