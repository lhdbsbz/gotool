package gotool

import "testing"

type Temp struct {
	name string
	age  int32
}

func TestIsNotEmpty(t *testing.T) {
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
			want: true,
		},
		{
			name: "string blank",
			args: args{
				" ",
			},
			want: false,
		},
		{
			name: "string empty",
			args: args{
				"",
			},
			want: false,
		},
		{
			name: "int 1",
			args: args{
				1,
			},
			want: true,
		},
		{
			name: "int 0",
			args: args{
				0,
			},
			want: false,
		},
		{
			name: "struct empty",
			args: args{
				Temp{},
			},
			want: false,
		},
		{
			name: "struct hasValue",
			args: args{
				Temp{
					name: "张三",
					age:  18,
				},
			},
			want: true,
		},
		{
			name: "ptr",
			args: args{
				&Temp{
					name: "张三",
					age:  18,
				},
			},
			want: true,
		},
		{
			name: "ptr empty",
			args: args{
				&Temp{},
			},
			want: true,
		},
		{
			name: "nil",
			args: args{
				nil,
			},
			want: false,
		},
		{
			name: "array",
			args: args{
				[2]int64{1, 2},
			},
			want: true,
		},
		{
			name: "array empty",
			args: args{
				[2]int64{},
			},
			want: false,
		},
		{
			name: "slice",
			args: args{
				[]int64{1, 2},
			},
			want: true,
		},
		{
			name: "slice empty",
			args: args{
				[]int64{},
			},
			want: false,
		},
		{
			name: "map",
			args: args{
				map[string]int64{
					"abc": 1,
				},
			},
			want: true,
		},
		{
			name: "map empty",
			args: args{
				make(map[string]int64),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotEmpty(tt.args.in); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
