package sort
/*
描述：插入排序（Insertion-Sort）的算法描述是一种简单直观的排序算法。它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
步骤:
	1）从第一个元素开始，该元素可以认为已经被排序；
	2)取出下一个元素，在已经排序的元素序列中从后向前扫描；
	3)如果该元素（已排序）大于新元素，将该元素移到下一位置；
	4)重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
	5)将新元素插入到该位置后；
	6)重复步骤2~5。
稳定性：稳定；时间复杂度：最好n,最坏n^2，平均n^2；空间复杂度：1
动图演示https://images2017.cnblogs.com/blog/849589/201710/849589-20171015225645277-1151100000.gif
*/
func InsertSort(array []int64) []int64 {
	size := len(array)
	for i := 1; i < size; i++ {
		current := array[i]
		for j := i - 1; j >= 0; j-- {
			if array[j] > current {
				array[j+1] = array[j]
				if j == 0 {
					array[0] = current
				}
			} else {
				array[j+1] = current
				break
			}
		}
	}
	return array
}
