package gotool

import "testing"

type Temp struct {
	name string
	age  int32
}

func TestIsEmpty(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string",
			args: args{
				" 1 ",
			},
			want: false,
		},
		{
			name: "string blank",
			args: args{
				" ",
			},
			want: true,
		},
		{
			name: "string empty",
			args: args{
				"",
			},
			want: true,
		},
		{
			name: "int 1",
			args: args{
				1,
			},
			want: false,
		},
		{
			name: "int 0",
			args: args{
				0,
			},
			want: true,
		},
		{
			name: "struct empty",
			args: args{
				Temp{},
			},
			want: true,
		},
		{
			name: "struct hasValue",
			args: args{
				Temp{
					name: "张三",
					age:  18,
				},
			},
			want: false,
		},
		{
			name: "ptr",
			args: args{
				&Temp{
					name: "张三",
					age:  18,
				},
			},
			want: false,
		},
		{
			name: "ptr empty",
			args: args{
				&Temp{},
			},
			want: false,
		},
		{
			name: "nil",
			args: args{
				nil,
			},
			want: true,
		},
		{
			name: "array",
			args: args{
				[2]int64{1, 2},
			},
			want: false,
		},
		{
			name: "array empty",
			args: args{
				[2]int64{},
			},
			want: true,
		},
		{
			name: "slice",
			args: args{
				[]int64{1, 2},
			},
			want: false,
		},
		{
			name: "slice empty",
			args: args{
				[]int64{},
			},
			want: true,
		},
		{
			name: "map",
			args: args{
				map[string]int64{
					"abc": 1,
				},
			},
			want: false,
		},
		{
			name: "map empty",
			args: args{
				make(map[string]int64),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.in); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
