package gotool

import (
	"testing"
	"time"
)

func TestToString(t *testing.T) {
	type args struct {
		item interface{}
	}
	now := time.Now()
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "时间转字符串",
			args: args{now},
			want: now.Format("2006-01-02 15:04:05"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.item); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMillisecond(t *testing.T) {
	type args struct {
		in time.Time
	}
	now := time.Now()
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "时间转毫秒",
			args: args{now},
			want: now.UnixNano() / 1e6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMillisecond(tt.args.in); got != tt.want {
				t.Errorf("ToMillisecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMillisecondToTime(t *testing.T) {
	now := time.Now()
	millisecond := ToMillisecond(now)
	toTime := MillToTime(millisecond)
	toMillisecond := ToMillisecond(toTime)
	if millisecond != toMillisecond {
		t.Errorf("MillToTime() = %v, want %v", toMillisecond, millisecond)
	}
}

func TestToSecond(t *testing.T) {
	now := time.Now()
	millisecond := ToMillisecond(now)
	toSecond := ToSecond(now)
	if toSecond != millisecond/1000 {
		t.Errorf("ToSecond() = %v, want %v", toSecond, millisecond/1000)
	}
}
