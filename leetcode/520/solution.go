package leetcode

import "unicode"

// 520. Detect capital
// Given a word, you need to judge whether the usage of capitals in it is right or not.
// We define the usage of capitals in a word to be right when one of the following cases holds:
// All letters in this word are capitals, like "USA".
// All letters in this word are not capitals, like "leetcode".
// Only the first letter in this word is capital, like "Google".
// Otherwise, we define that this word doesn't use capitals in a right way.

func detectCapitalUse(word string) bool {
	if IsUpper(word) || IsLower(word) || firstUpper(word) {
		return true
	}
	return false
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func firstUpper(s string) bool {
	for i, r := range s {
		if i == 0 {
			if !unicode.IsUpper(r) && unicode.IsLetter(r) {
				return false
			}
		} else {
			if !unicode.IsLower(r) && unicode.IsLetter(r) {
				return false
			}
		}
	}
	return true
}
