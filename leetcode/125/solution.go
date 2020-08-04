package leetcode

// 125. Valid Palindrome
// Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
// Note: For the purpose of this problem, we define empty string as valid palindrome.

import "strings"

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < len(s); {
		if !isAlphanumeric(s[i]) {
			i++
			continue
		}

		if !isAlphanumeric(s[j]) {
			j--
			continue
		}

		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
