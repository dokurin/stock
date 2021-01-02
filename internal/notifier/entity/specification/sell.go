package specification

import "github.com/dokurin/stock/internal/notifier/entity/model"

type Sell struct {
	tradeCondition model.TradeCondition
}

func NewSell(cond model.TradeCondition) *Sell {
	return &Sell{
		tradeCondition: cond,
	}
}

func (s *Sell) ShouldTrade(d model.StockDifference) bool {
	return d.Ratio() <= s.tradeCondition.SellRatio
}
