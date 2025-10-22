package leetcode

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	var (
		m = make(map[string][]string)
	)
	for _, str := range strs {
		var (
			key = sortString(str)
		)
		if _, ok := m[key]; !ok {
			m[key] = make([]string, 0)
			m[key] = append(m[key], str)
			continue
		}

		m[key] = append(m[key], str)
	}

	var (
		result [][]string
	)
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func sortString(s string) string {
	var (
		runes = []rune(s)
	)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
