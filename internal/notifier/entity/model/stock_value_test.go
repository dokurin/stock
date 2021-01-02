package model_test

import (
	"testing"
	"time"

	. "github.com/dokurin/stock/internal/notifier/entity/model"
	"github.com/stretchr/testify/assert"
)

func TestStockValue_IsBlank(t *testing.T) {
	type fields struct {
		Open  float64
		Low   float64
		High  float64
		Close float64
		Time  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "株式情報が空",
			fields: fields{},
			want:   true,
		},
		{
			name: "株式情報が空ではない",
			fields: fields{
				Open:  1,
				Low:   0,
				High:  0,
				Close: 0,
				Time:  time.Time{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StockValue{
				Open:  tt.fields.Open,
				Low:   tt.fields.Low,
				High:  tt.fields.High,
				Close: tt.fields.Close,
				Time:  tt.fields.Time,
			}
			got := s.IsBlank()
			assert.Equal(t, tt.want, got)
		})
	}
}
