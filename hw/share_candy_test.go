package huawei

import "testing"

func TestShareCandy(t *testing.T) {
	type args struct {
		candy int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1:1",
			args: args{
				candy: 1,
			},
			want: 0,
		},
		{
			name: "test1:2",
			args: args{
				candy: 2,
			},
			want: 1,
		},
		{
			name: "test1:3",
			args: args{
				candy: 3,
			},
			want: 2,
		},
		{
			name: "test1:15",
			args: args{
				candy: 15,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShareCandy(tt.args.candy); got != tt.want {
				t.Errorf("ShareCandy() = %v, want %v", got, tt.want)
			}
		})
	}
}
