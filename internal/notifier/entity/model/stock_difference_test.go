package model_test

import (
	"testing"

	"github.com/benbjohnson/clock"
	. "github.com/dokurin/stock/internal/notifier/entity/model"
	"github.com/stretchr/testify/assert"
)

func TestNewStockDifference(t *testing.T) {
	type args struct {
		brand   Brand
		before  StockValue
		current StockValue
	}
	tests := []struct {
		name string
		args args
		want *StockDifference
	}{
		{
			name: "new stock difference",
			args: args{
				brand: "test_brand",
				before: StockValue{
					Open:  1.1,
					Low:   2.2,
					High:  3.3,
					Close: 4.4,
					Time:  clock.NewMock().Now(),
				},
				current: StockValue{
					Open:  2.2,
					Low:   3.3,
					High:  4.4,
					Close: 5.5,
					Time:  clock.NewMock().Now(),
				},
			},
			want: new(StockDifference).
				WithBrand("test_brand").
				WithBefore(StockValue{
					Open:  1.1,
					Low:   2.2,
					High:  3.3,
					Close: 4.4,
					Time:  clock.NewMock().Now(),
				}).
				WithCurrent(StockValue{
					Open:  2.2,
					Low:   3.3,
					High:  4.4,
					Close: 5.5,
					Time:  clock.NewMock().Now(),
				}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewStockDifference(tt.args.brand, tt.args.before, tt.args.current)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStockDifference_Ratio(t *testing.T) {
	type fields struct {
		before  StockValue
		current StockValue
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "株価2.0から1.5に下落している場合は, 0.75を返す",
			fields: fields{
				before: StockValue{
					Open:  0,
					Low:   0,
					High:  0,
					Close: 2.0,
					Time:  clock.NewMock().Now(),
				},
				current: StockValue{
					Open:  0,
					Low:   0,
					High:  0,
					Close: 1.5,
					Time:  clock.NewMock().Now(),
				},
			},
			want: 0.75,
		},
		{
			name: "株価2.0から2.5に上昇している場合は, 1.25を返す",
			fields: fields{
				before: StockValue{
					Open:  0,
					Low:   0,
					High:  0,
					Close: 2.0,
					Time:  clock.NewMock().Now(),
				},
				current: StockValue{
					Open:  0,
					Low:   0,
					High:  0,
					Close: 2.5,
					Time:  clock.NewMock().Now(),
				},
			},
			want: 1.25,
		},
		{
			name: "前回の株価が存在しない場合は、1.0を返す",
			fields: fields{
				before: StockValue{
					Open:  0,
					Low:   0,
					High:  0,
					Close: 0,
					Time:  clock.NewMock().Now(),
				},
				current: StockValue{
					Open:  0,
					Low:   0,
					High:  0,
					Close: 2.5,
					Time:  clock.NewMock().Now(),
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewStockDifference("test_brand", tt.fields.before, tt.fields.current).Ratio()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStockDifference_Margin(t *testing.T) {
	type fields struct {
		before  StockValue
		current StockValue
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "前回1000, 現在1200の場合",
			fields: fields{
				before: StockValue{
					Open:  0,
					Low:   0,
					High:  0,
					Close: 1000,
					Time:  clock.NewMock().Now(),
				},
				current: StockValue{
					Open:  0,
					Low:   0,
					High:  0,
					Close: 1200,
					Time:  clock.NewMock().Now(),
				},
			},
			want: 200,
		},
		{
			name: "前回1000, 現在800の場合",
			fields: fields{
				before: StockValue{
					Open:  0,
					Low:   0,
					High:  0,
					Close: 1000,
					Time:  clock.NewMock().Now(),
				},
				current: StockValue{
					Open:  0,
					Low:   0,
					High:  0,
					Close: 800,
					Time:  clock.NewMock().Now(),
				},
			},
			want: -200,
		},
		{
			name: "前回の株式情報が空の場合",
			fields: fields{
				before: StockValue{},
				current: StockValue{
					Open:  0,
					Low:   0,
					High:  0,
					Close: 800,
					Time:  clock.NewMock().Now(),
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewStockDifference("test_brand", tt.fields.before, tt.fields.current).Margin()
			assert.Equal(t, tt.want, got)
		})
	}
}
