package huawei

import "testing"

func TestStringQuack(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1 quack",
			args: args{
				in: "quack",
			},
			want: 1,
		},
		{
			name: "test2 quackquack",
			args: args{
				in: "quackquack",
			},
			want: 1,
		},
		{
			name: "test3 quqackuack",
			args: args{
				in: "quqackuack",
			},
			want: 2,
		},
		{
			name: "test4 quackquook",
			args: args{
				in: "quackquook",
			},
			want: -1,
		},
		{
			name: "test5 quqackuackquack",
			args: args{
				in: "quqackuackquack",
			},
			want: 2,
		},
		{
			name: "test6 qququaauqccauqkkcauqqkcauuqkcaaukccakkck",
			args: args{
				in: "qququaauqccauqkkcauqqkcauuqkcaaukccakkck",
			},
			want: 8,
		},

		{
			name: "test7 ququackackqakcu",
			args: args{
				in: "ququackackqakcu",
			},
			want: 2,
		},
		{
			name: "test8 quacqkuac",
			args: args{
				in: "quacqkuac",
			},
			want: 1,
		},
		{
			name: "test9 quacquackquack",
			args: args{
				in: "quacquackquack",
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringQuack(tt.args.in); got != tt.want {
				t.Errorf("StringQuack() = %v, want %v", got, tt.want)
			}
		})
	}
}
