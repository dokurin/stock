package model_test

import (
	"testing"

	"github.com/benbjohnson/clock"
	. "github.com/dokurin/stock/internal/notifier/entity/model"
	"github.com/stretchr/testify/assert"
)

func TestStockHistoryFromStock(t *testing.T) {
	type args struct {
		id    StockHistoryID
		brand Brand
		value StockValue
	}
	tests := []struct {
		name string
		args args
		want *StockHistory
	}{
		{
			name: "new history",
			args: args{
				id:    1,
				brand: "test_brand",
				value: StockValue{
					Open:  1.1,
					Low:   2.2,
					High:  3.3,
					Close: 4.4,
					Time:  clock.NewMock().Now(),
				},
			},
			want: &StockHistory{
				ID:    1,
				Brand: "test_brand",
				Value: StockValue{
					Open:  1.1,
					Low:   2.2,
					High:  3.3,
					Close: 4.4,
					Time:  clock.NewMock().Now(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewStockHistory(tt.args.id, tt.args.brand, tt.args.value)
			assert.Equal(t, tt.want, got)
		})
	}
}
