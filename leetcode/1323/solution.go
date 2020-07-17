package leetcode

import (
	"strconv"
	"strings"
)

// 1323. Maximum 69 Number
// Given a positive integer num consisting only of digits 6 and 9.
// Return the maximum number you can get by changing at most one digit (6 becomes 9, and 9 becomes 6).

func maximum69Number(num int) int {
	str := strconv.Itoa(num)
	str = strings.Replace(str, "6", "9", 1)
	result, _ := strconv.Atoi(str)
	return result
}
