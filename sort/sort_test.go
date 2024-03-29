package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		arrs []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "test1",
			args: args{
				[]int64{1, 2, 4, 5, 7, 6, 8, 3, 9, 0},
			},
			want: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.arrs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestSelectSort(t *testing.T) {
	type args struct {
		arrs []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "test1",
			args: args{
				[]int64{1, 2, 4, 5, 7, 6, 8, 3, 9, 0},
			},
			want: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectSort(tt.args.arrs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestInsertSort(t *testing.T) {
	type args struct {
		arrs []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "test1",
			args: args{
				[]int64{1, 2, 4, 5, 7, 6, 8, 3, 9, 0},
			},
			want: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertSort(tt.args.arrs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
