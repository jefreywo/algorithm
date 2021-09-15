package sort

// 合并两个数组，并按从小到大对合并好的数组进行排序
// arr1 和 arr2 均已按由小到大做好排序
func mergeArrayAnd(arr1 []int64, arr2 []int64) []int64 {
	index1 := len(arr1) - 1
	index2 := len(arr2) - 1

	arr1 = append(arr1, arr2...)

	for index1 >= 0 || index2 >= 0 {
		if index1 < 0 {
			arr1[index1+index2+1] = arr2[index2]
			index2--
			continue
		}
		if index2 < 0 {
			arr1[index1+index2+1] = arr1[index1]
			index1--
			continue
		}
		if arr1[index1] >= arr2[index2] {
			arr1[index1+index2+1] = arr1[index1]
			index1--
		} else {
			arr1[index1+index2+1] = arr2[index2]
			index2--
		}
	}
	return arr1
}
