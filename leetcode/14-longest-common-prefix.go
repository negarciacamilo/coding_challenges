/*
	Write a function to find the longest common prefix string amongst an array of strings.

	If there is no common prefix, return an empty string "".

	 

	Example 1:

	Input: strs = ["flower","flow","flight"]
	Output: "fl"
	Example 2:

	Input: strs = ["dog","racecar","car"]
	Output: ""
	Explanation: There is no common prefix among the input strings.
	 

	Constraints:

	1 <= strs.length <= 200
	0 <= strs[i].length <= 200
	strs[i] consists of only lower-case English letters.
*/

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