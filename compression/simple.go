package compression

import "fmt"

func SimpleCompression(source string) string {
	size := len(source)
	if size < 2 {
		return source
	}

	ret := ""
	tmp := source[0]
	sum := 1

	fn := func(s byte) {
		if sum > 1 {

			ret += fmt.Sprintf("%d%s", sum, string(s))
		} else {
			ret += string(s)
		}
	}

	for i := 1; i <= size; i++ {
		if i == size {
			fn(tmp)
			break
		}
		if source[i] == tmp {
			sum++
			continue
		}

		fn(tmp)
		tmp = source[i]
		sum = 1
	}
	return ret
}
