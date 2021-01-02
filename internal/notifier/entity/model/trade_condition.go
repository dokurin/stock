package model

type Brand string

type TradeCondition struct {
	Brand     Brand
	SellRatio float64
	BuyPrice  float64
}
