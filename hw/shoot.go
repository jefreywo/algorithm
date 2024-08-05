package huawei

func ShootMain() {

}

func Shoot(n int, ids []int, score []int) []int {
	var group = make(map[int][]int)
	for i, id := range ids {
		if _, ok := group[id]; !ok {
			group[id] = []int{score[i]}
			continue
		}
		group[id] = append(group[id], score[i])
	}
	maxSum := func(nums []int) int {
		if len(nums) < 3 {
			return -1
		}
		var maxa = nums[0]
		var maxb = nums[1]
		var maxc = nums[2]
		if len(nums) > 3 {
			for i := 3; i < len(nums); i++ {
				v := nums[i]
				if v > maxa {
					maxa = v
					continue
				}
				if v > maxb {
					maxb = v
					continue
				}
				if v > maxc {
					maxa = v
					continue
				}
			}
		}

		return maxa + maxb + maxc
	}
	var resultMap = make(map[int]int)
	var scoreArray = make([]int, 0, len(group))
	for id, scores := range group {
		scoreSum := maxSum(scores)
		resultMap[scoreSum] = id
		scoreArray = append(scoreArray, scoreSum)
	}

	var result []int
	return result
}

/*
给定一个射击比赛成绩单，包含多个选手若干次射击的成绩分数，请对每个选手按其最高三个分数之

和进行降序排名，输出降序排名后的选手id序列。

题目解析：

给一个数字表示射击的次数，然后给几个选手进行（乱序）射击，生成对应的成绩！

条件如下：

一个选手可以有多个射击成绩的分数，且次序不固定

如果一个选手成绩少于3个，则认为选手的所有成绩无效，排名忽略该选手

如果选手的成绩之和相等，则相等的选手按照其id降序排列

输入描述：

输入第一行，一个整数N，表示该场比赛总共进行了N次射击，产生N个成绩分数（2<=N<=100）

输入第二行，一个长度为N整数序列，表示参与每次射击的选手id（0<=id<=99）

输入第三行，一个长度为N整数序列，表示参与每次射击选手对应的成绩（0<=成绩<=100）

输出描述：

符合题设条件的降序排名后的选手ID序列

示例

输入：

13

3,3,7,4,4,4,4,7,7,3,5,5,5

53,80,68,24,39,76,66,16,100,55,53,80,55

输出：

5,3,7,4

*/
