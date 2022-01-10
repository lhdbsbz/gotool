package gotool

import "testing"

func TestFloatMoneyEquals(t *testing.T) {
	type args struct {
		f1        float64
		f2        float64
		precision []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "测试",
			args: args{
				f1:        1.099,
				f2:        1.10,
				precision: []int{2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatMoneyEquals(tt.args.f1, tt.args.f2, tt.args.precision...); got != tt.want {
				t.Errorf("FloatMoneyEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}
