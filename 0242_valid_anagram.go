package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var (
		m = make(map[string]int)
	)
	for _, ch := range s {
		m[string(ch)]++
	}

	for _, ch := range t {
		m[string(ch)]--

		if m[string(ch)] < 0 {
			return false
		}
	}

	return true
}
