package huawei

import "testing"

func TestJobNumber(t *testing.T) {
	type args struct {
		people int
		letter int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testA,1,1",
			args: args{
				people: 1,
				letter: 1,
			},
			want: 1,
		},
		{
			name: "testB,26,1",
			args: args{
				people: 26,
				letter: 1,
			},
			want: 1,
		},
		{
			name: "testC,2600,1",
			args: args{
				people: 2600,
				letter: 1,
			},
			want: 2,
		},
		{
			name: "testD,123456,3",
			args: args{
				people: 123456,
				letter: 3,
			},
			want: 1,
		},
		{
			name: "testE,28256,2",
			args: args{
				people: 28256,
				letter: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JobNumber(tt.args.people, tt.args.letter); got != tt.want {
				t.Errorf("JobNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
