package model

type StockDifference struct {
	Brand   Brand
	before  StockValue
	current StockValue
}

func NewStockDifference(brand Brand, before, current StockValue) *StockDifference {
	return &StockDifference{
		Brand:   brand,
		before:  before,
		current: current,
	}
}

func (s StockDifference) Ratio() float64 {
	if s.before.IsBlank() {
		return 1
	}
	return s.current.Close / s.before.Close
}

func (s StockDifference) Margin() float64 {
	if s.before.IsBlank() {
		return 0
	}
	return s.current.Close - s.before.Close
}
