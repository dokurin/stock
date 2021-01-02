package usecase

import (
	"context"

	"github.com/dokurin/stock/internal/notifier/entity/model"
)

type (
	TradeCheckerInputPort interface {
	}

	TradeCheckerOutputPort interface {
	}
)

type (
	stockHistoryRepository interface {
		GetLatest(context.Context) (*model.StockHistory, error)
		Store(context.Context, model.StockHistory) error
	}
	tradeConditionRepository interface {
		GetAll(context.Context) ([]*model.TradeCondition, error)
	}
	stockService interface {
		FetchCurrent(context.Context, model.Brand) (*model.Stock, error)
	}
)

type TradeChecker struct {
	repository struct {
		stockHistory stockHistoryRepository
	}
	service struct {
		stock stockService
	}
}

func (t *TradeChecker) Check(ctx context.Context) {
	_, err := t.repository.stockHistory.GetLatest(ctx)
	if err != nil {
		panic(err)
	}
	t.service.stock.FetchCurrent(ctx, "")
}
