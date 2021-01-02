package repository

import (
	"context"

	"github.com/dokurin/stock/internal/notifier/entity/model"
)

type StockHistory interface {
	GetLatest(context.Context, model.Brand) (*model.StockHistory, error)
}
