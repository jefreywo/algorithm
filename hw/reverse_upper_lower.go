package huawei

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseMain() {
	var k int
	var source string
	fmt.Scanf("%d", &k)
	fmt.Scanf("%s", &source)
	fmt.Printf("输入k:%d, source:%s\n", k, source)
	fmt.Println(ReverseUpperLower(source, k))
}

func ReverseUpperLower(source string, k int) string {
	sourceArray := strings.Split(source, "-")
	if len(sourceArray) <= 1 {
		return source
	}
	headStr := sourceArray[0]

	newSource := strings.Join(sourceArray[1:], "")

	newLen := len(newSource)

	minIndex := 0
	var result = []string{headStr}

	handle := func(s string) string {
		fmt.Println("---handle s:", s)
		var upperNum, lowerNum int
		for i := 0; i < len(s); i++ {

			if unicode.IsUpper(rune(s[i])) {
				upperNum++
			}
			if unicode.IsLower(rune(s[i])) {
				lowerNum++
			}
		}
		if upperNum > lowerNum {
			return strings.ToUpper(s)
		} else if upperNum < lowerNum {
			return strings.ToLower(s)
		}
		return s
	}

	var subStr string
	for i := 1; i <= newLen; i++ {
		if i%k == 0 {
			subStr = newSource[minIndex:i]
			result = append(result, handle(subStr))
			minIndex = i
		}
	}
	if newLen%k > 0 {
		result = append(result, handle(newSource[minIndex:]))
	}
	return strings.Join(result, "-")
}

/*
给定一个非空字符串S，其被N个‘-’分隔成N+1的子串，给定正整数K，要求除第一个子串外，其余的子串每K个字符组成新的子串，并用‘-’分隔。对于新组成的每一个子串，如果它含有的小写字母比大写字母多，则将这个子串的所有大写字母转换为小写字母；反之，如果它含有的大写字母比小写字母多，则将这个子串的所有小写字母转换为大写字母；大小写字母的数量相等时，不做转换。
输入描述:
输入为两行，第一行为参数K，第二行为字符串S。
输出描述:
输出转换后的字符串。
示例1
输入
3
12abc-abCABc-4aB@
输出
12abc-abc-ABC-4aB-@
说明
子串为12abc、abCABc、4aB@，第一个子串保留，后面的子串每3个字符一组为abC、ABc、4aB、@，abC中小写字母较多，转换为abc，ABc中大写字母较多，转换为ABC，4aB中大小写字母都为1个，不做转换，@中没有字母，连起来即12abc-abc-ABC-4aB-@
示例2
输入
12
12abc-abCABc-4aB@
输出
12abc-abCABc4aB@
说明
子串为12abc、abCABc、4aB@，第一个子串保留，后面的子串每12个字符一组为abCABc4aB@，这个子串中大小写字母都为4个，不做转换，连起来即12abc-abCABc4aB@
*/
