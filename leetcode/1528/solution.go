package leetcode

import "strings"

// 1528. Shuffle String
// Given a string s and an integer array indices of the same length.
// The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.
// Return the shuffled string.

func restoreString(s string, indices []int) string {
	a := make([]string, len(indices))

	for i, j := range indices {
		a[j] = string(s[i])
	}

	return strings.Join(a[:], "")

}
