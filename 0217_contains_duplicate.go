package leetcode

func containsDuplicate(nums []int) bool {
	var (
		m = make(map[int]int)
	)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		} else {
			m[num] = 1
		}
	}

	return false
}
