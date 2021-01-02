package model

func (s *StockDifference) WithBefore(b StockValue) *StockDifference {
	s.before = b
	return s
}

func (s *StockDifference) WithCurrent(c StockValue) *StockDifference {
	s.current = c
	return s
}

func (s *StockDifference) WithBrand(b Brand) *StockDifference {
	s.Brand = b
	return s
}
