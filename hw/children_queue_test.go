package huawei

import (
	"reflect"
	"testing"
)

func TestChildrenQueue(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1 ",
			args: args{
				input: "4 3 5 7 8",
			},
			want: "4 3 7 5 8",
		},
		{
			name: "test2 ",
			args: args{
				input: "4 1 3 5 2",
			},
			want: "4 1 5 2 3",
		},
		{
			name: "test3 1 1 1 1 1",
			args: args{
				input: "1 1 1 1 1",
			},
			want: "1 1 1 1 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChildrenQueue(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChildrenQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
