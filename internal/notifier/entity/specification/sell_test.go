package specification_test

import (
	"testing"

	"github.com/benbjohnson/clock"
	"github.com/dokurin/stock/internal/notifier/entity/model"
	. "github.com/dokurin/stock/internal/notifier/entity/specification"
	"github.com/stretchr/testify/assert"
)

func TestSell_ShouldTrade(t *testing.T) {
	type args struct {
		d model.StockDifference
	}
	tests := []struct {
		name string
		s    *Sell
		args args
		want bool
	}{
		{
			name: "SellRatio0.97_3%(100 -> 97)の下落",
			s: NewSell(model.TradeCondition{
				Brand:     "test_brand",
				SellRatio: 0.97,
			}),
			args: args{
				d: *model.NewStockDifference(
					"test_brand",
					model.StockValue{
						Close: 100,
						Time:  clock.NewMock().Now(),
					}, model.StockValue{
						Close: 97,
						Time:  clock.NewMock().Now(),
					},
				),
			},
			want: true,
		},
		{
			name: "SellRatio0.97_2%(100 -> 98)の下落",
			s: NewSell(model.TradeCondition{
				Brand:     "test_brand",
				SellRatio: 0.97,
			}),
			args: args{
				d: *model.NewStockDifference(
					"test_brand",
					model.StockValue{
						Close: 100,
						Time:  clock.NewMock().Now(),
					}, model.StockValue{
						Close: 98,
						Time:  clock.NewMock().Now(),
					},
				),
			},
			want: false,
		},
		{
			name: "新規の株式(前の株式情報が空欄)",
			s: NewSell(model.TradeCondition{
				Brand:     "test_brand",
				SellRatio: 0.97,
			}),
			args: args{
				d: *model.NewStockDifference(
					"test_brand",
					model.StockValue{},
					model.StockValue{
						Close: 100,
						Time:  clock.NewMock().Now(),
					},
				),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.s.ShouldTrade(tt.args.d))
		})
	}
}
