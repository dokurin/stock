package model

import "time"

type StockValue struct {
	Open  float64
	Low   float64
	High  float64
	Close float64
	Time  time.Time
}

func (s StockValue) IsBlank() bool {
	return s.Open == 0 && s.Low == 0 && s.High == 0 && s.Close == 0
}
