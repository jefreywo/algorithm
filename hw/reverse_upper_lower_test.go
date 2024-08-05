package huawei

import "testing"

func TestReverseUpperLower(t *testing.T) {
	type args struct {
		source string
		k      int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				source: "12abc-abCABc-4aB@",
				k:      3,
			},
			want: "12abc-abc-ABC-4aB-@",
		},
		{
			name: "test2",
			args: args{
				source: "12abc-abCABc-4aB@",
				k:      12,
			},
			want: "12abc-abCABc4aB@",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseUpperLower(tt.args.source, tt.args.k); got != tt.want {
				t.Errorf("ReverseUpperLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
