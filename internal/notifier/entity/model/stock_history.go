package model

type StockHistoryID int

type StockHistory struct {
	ID    StockHistoryID
	Brand Brand
	Value StockValue
}

func NewStockHistory(id StockHistoryID, brand Brand, value StockValue) *StockHistory {
	return &StockHistory{
		ID:    id,
		Brand: brand,
		Value: value,
	}
}
