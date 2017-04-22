package main

type Conversion struct {
	ConvertFrom string `json:"from"`
	ConvertTo   string `json:"to"`
}

type Conversions []Conversion

type Currency struct {
	Coin  string  `json: "coin"`
	Value float64 `json: "value"`
}

type Currencies []Currency
