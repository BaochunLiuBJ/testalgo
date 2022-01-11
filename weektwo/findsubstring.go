package weektwo


func findSubstring(s string, words []string) []int {
	wordsNum := len(words)
	result := make([]int, 0)
	if wordsNum == 0 {
		return result
	}
	wordLen := len(words[0])

	wordsMap := make(map[string]int)
	for _, word := range words {
		count, ok := wordsMap[word]
		if !ok {
			wordsMap[word] = 1
		} else {
			wordsMap[word] = count + 1
		}
	}

	wl := wordsNum * wordLen

	for i := 0; i+wl <= len(s); i++ {
		st := []byte(s)[i : i+wl]
		tMap := make(map[string]int)
		for j:= 0 ; j + wordLen <= wl; j=j+wordLen {
			t := st[j:j+wordLen]
			c, ok := tMap[string(t)]
			if !ok {
				tMap[string(t)] = 1
			} else {
				tMap[string(t)] = c + 1 
			}
		}
		if equalMap(wordsMap, tMap ) {
			result = append(result, i)
		}
	}
	return result  
}

func equalMap(wMap map[string]int, tMap map[string]int) bool {
	for k, v := range wMap {
		c, ok := tMap[k]
		if ok && v == c {
			continue
		} else {
			return false 
		}
	}
	
	for k, v := range tMap {
		c, ok := wMap[k]
		if !ok || v != c {
			return false 
		}
	}

	return true 
}