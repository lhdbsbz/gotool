package gotool

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestCeilDecimal(t *testing.T) {
	type args struct {
		amount    decimal.Decimal
		precision int32
	}
	tests := []struct {
		name string
		args args
		want decimal.Decimal
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				amount:    decimal.NewFromFloat(1.8423),
				precision: 0,
			},
			want: decimal.NewFromFloat(2.0),
		},
		{
			name: "1",
			args: args{
				amount:    decimal.NewFromFloat(1.8423),
				precision: 1,
			},
			want: decimal.NewFromFloat(1.9),
		},
		{
			name: "2",
			args: args{
				amount:    decimal.NewFromFloat(1.8423),
				precision: 2,
			},
			want: decimal.NewFromFloat(1.85),
		},
		{
			name: "3",
			args: args{
				amount:    decimal.NewFromFloat(1.8423),
				precision: 3,
			},
			want: decimal.NewFromFloat(1.843),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CeilDecimal(tt.args.amount, tt.args.precision); !got.Equal(tt.want) {
				t.Errorf("CeilDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloorDecimal(t *testing.T) {
	type args struct {
		amount    decimal.Decimal
		precision int32
	}
	tests := []struct {
		name string
		args args
		want decimal.Decimal
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				amount:    decimal.NewFromFloat(1.8423),
				precision: 0,
			},
			want: decimal.NewFromFloat(1.0),
		},
		{
			name: "1",
			args: args{
				amount:    decimal.NewFromFloat(1.8423),
				precision: 1,
			},
			want: decimal.NewFromFloat(1.8),
		},
		{
			name: "2",
			args: args{
				amount:    decimal.NewFromFloat(1.8423),
				precision: 2,
			},
			want: decimal.NewFromFloat(1.84),
		},
		{
			name: "3",
			args: args{
				amount:    decimal.NewFromFloat(1.8423),
				precision: 3,
			},
			want: decimal.NewFromFloat(1.842),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloorDecimal(tt.args.amount, tt.args.precision); !got.Equal(tt.want) {
				t.Errorf("CeilDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
