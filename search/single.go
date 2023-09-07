package search

func SequenceSearch(source []int64, t int64) int {
	for i := 0; i < len(source); i++ {
		if source[i] == t {
			return i
		}
	}
	return -1
}

func Search(source []int64, t int64) int {
	var low, high, mid = 0, len(source) - 1, 0
	for low <= high {
		mid = low + (low+high)/2
		if source[mid] == t {
			return mid
		}
		if source[mid] < t {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func recursionSearch(source []int64, t int64) int {
	if len(source) == 0 {
		return -1
	}
	if len(source) == 1 && source[0] == t {
		return 0
	}
	mid := len(source) / 2
	if source[mid] > t {
		return recursionSearch(source[:mid-1], t)
	}

	return recursionSearch(source[mid+1:], t)
}
