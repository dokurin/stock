package repository

import (
	"context"

	"github.com/dokurin/stock/internal/notifier/entity/model"
)

type TradeCondition interface {
	Get(context.Context, model.Brand) (*model.TradeCondition, error)
}
