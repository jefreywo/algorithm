package sort

import (
	"fmt"
	"testing"
)

func Test_mergeArrayAnd(t *testing.T) {
	var arr1 = []int64{2, 4, 6, 8, 10, 20, 40, 60, 80, 100}
	var arr2 = []int64{1, 3, 5, 7, 9, 11, 30, 50, 70, 90}
	fmt.Println(mergeArrayAnd(arr1, arr2))
}
