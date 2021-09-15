package sort

import (
	"log"
)

/*
//描述：冒泡排序重复地走访过要排序地数列，一次比较两个元素，如果它们地顺序错误就交换过来。走访数列地工作会重复地进行直到没有再需要交换。
//步骤:
// 1）比较相邻地元素，如果第一个比第二个大，就交换它们，
// 2）对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
// 3)针对所有的元素重复以上的步骤，除了最后一个；
// 4)重复步骤1~3，直到排序完成。
//稳定性：稳定；时间复杂度：最好n,最坏n^2，平均n^2；空间复杂度：1，
*/
func BubbleSort(arrs []int64) []int64 {
	log.Println("Input:", arrs)
	var size = len(arrs)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if arrs[j] > arrs[j+1] {
				arrs[j], arrs[j+1] = arrs[j+1], arrs[j]
			}
		}
	}
	log.Println("Output:", arrs)
	return arrs
}
