package leetcode

//242. Valid Anagram
//Given two strings s and t , write a function to determine if t is an anagram of s.

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, j := range s {
		m[j]++
	}

	for _, x := range t {
		m[x]--
	}

	for k := range m {
		if m[k] != 0 {
			return false
		}
	}
	return true
}
