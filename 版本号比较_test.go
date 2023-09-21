package gotool

import "testing"

func TestVersionA_BiggerThanOrEqual_VersionB(t *testing.T) {
	type args struct {
		versionA string
		versionB string
	}
	var tests = []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				versionA: "1.0.0",
				versionB: "1.0.0.1",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				versionA: "1.0.0",
				versionB: "1.0.0",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				versionA: "1.0.0.1",
				versionB: "1.0.0",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				versionA: "0.0.0.1",
				versionB: "1.0.0",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VersionA_BiggerThanOrEqual_VersionB(tt.args.versionA, tt.args.versionB); got != tt.want {
				t.Errorf("VersionA_BiggerThanOrEqual_VersionB() = %v, want %v", got, tt.want)
			}
		})
	}
}
