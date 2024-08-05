package huawei

import "fmt"

/*
一群大雁往南 飞，给定一个字符串记录地面上的游客听到的大雁叫声，请给出叫声最少由几只大雁发出。

具体的:

1.大雁发出的完整叫声为"quack"，因为有多只大雁同一时间嘎嘎作响，所以字符串中可能会混合多个"quack"。

2.大雁会依次完整发出"quack"，即字符串中’q’ ,‘u’, ‘a’, ‘c’, ‘k’ 这5个字母按顺序完整存在才能计数为一只大雁。如果不完整或者没有按顺序则不予计数。

3.如果字符串不是由’q’, ‘u’, ‘a’, ‘c’, ‘k’ 字符组合而成，或者没有找到一只大雁，请返回-1。

输入描述

一个字符串，包含大雁quack的叫声。1 <= 字符串长度 <= 1000，字符串中的字符只有’q’, ‘u’, ‘a’, ‘c’, ‘k’

输出描述

大雁的数量

示例1：

输入：quackquack

输出：1
一只大雁叫了两声

示例2：

输入：quqackuack

输出：2
至少两只大雁各叫了一声quack

示例3：

输入：quackquook

输出：-1
给出的字符串不是quack的有效组合

实例3
输入：quqackuackquack
输出：2
*/
// q u a c k
func StringQuack(in string) int {
	var q, k []int
	var countU, countA, countC int

	for i, v := range in {
		switch string(v) {
		case "q":
			q = append(q, i)

		case "u":
			if countU < len(q) {
				countU++
			}

		case "a":
			if countA < countU {
				countA++
			}

		case "c":
			if countC < countA {
				countC++
			}

		case "k":
			if countC > 0 {
				k = append(k, i)
				countU--
				countA--
				countC--
			}

		default:
			return -1
		}
		//fmt.Println("Switch countQ, countU, countA, countC, countK", string(v), q, countU, countA, countC, k)
	}

	var size = len(q)
	if len(k) < size {
		size = len(k)
	}
	var res int
	var flag bool
	for i := 0; i < size; i++ {
		if q[i]+4 == k[i] {
			flag = true
			continue
		}
		res++
	}

	if flag && res == 0 {
		res++
	}

	if res == 0 {
		return -1
	}
	fmt.Println("countQ, countU, countA, countC, countK", q, countU, countA, countC, k, res)
	return res
}