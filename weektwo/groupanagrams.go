package weektwo

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		t := []byte(s)
		sort.SliceStable(t, func(i, j int) bool {
			return t[i] < t[j]
		})
		if strSlice, ok := m[string(t)]; ok {
			strSlice = append(strSlice, s)
			m[string(t)] = strSlice
		} else {
			if len(strSlice) == 0 {
				strSlice = make([]string, 0)
			}
			strSlice = append(strSlice, s)
			m[string(t)] = strSlice
		}
	}
	result := make([][]string, 0)
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
