package leetcode

// 344. Reverse String
// Write a function that reverses a string. The input string is given as an array of characters char[].
// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
// You may assume all the characters consist of printable ascii characters.

func reverseString(s []byte) string {
	bytes := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	return string(bytes)
}
