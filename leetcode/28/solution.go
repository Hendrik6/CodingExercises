package leetcode

// 28. Implement strStr()
// Implement strStr().
// Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

func strStr(haystack string, needle string) int {
	lengthSub := len(needle)
	if lengthSub == 0 {
		return 0
	}
	length := len(haystack)
	diffLength := length - lengthSub
	for i := 0; i <= diffLength; i++ {
		if haystack[i:i+lengthSub] == needle {
			return i
		}
	}
	return -1
}
