func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
        return strs[0]
    }
    
	prefixes := make(map[string]int)
	for _, s := range strs {
		for i := 0; i <= len(s); i++ {
			p := fmt.Sprintf("%s", s[:i])
			if _, ok := prefixes[p]; ok && p != "" {
				prefixes[p]++
			} else {
				prefixes[p] = 1
			}
		}
	}

	max := 0
	var p []string
	for k, v := range prefixes {
		if v == 1 {
			continue
		}
		if v >= max {
			if v == max {
				p = append(p, k)
			} else {
				p = []string{k}
			}
			max = v
		}
	}

	if len(p) == 0 {
		return ""
	}
    
    if max != len(strs) {
        return ""
    }

	maxLen := 0
	maxIndex := 0
	for i, v := range p {
		if len(v) > maxLen {
			maxIndex = i
			maxLen = len(v)
		}
	}
	return p[maxIndex]
}