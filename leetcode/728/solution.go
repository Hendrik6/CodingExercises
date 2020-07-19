package leetcode

// 728. Self Dividing Numbers
// A self-dividing number is a number that is divisible by every digit it contains.
// For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.
// Also, a self-dividing number is not allowed to contain the digit zero.
// Given a lower and upper number bound, output a list of every possible self dividing number, including the bounds if possible.

func selfDividingNumbers(left int, right int) []int {
	ret := make([]int, 0, right-left+1)
	for i := left; i <= right; i++ {
		if isValid(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

func isValid(num int) bool {
	for n := num; n > 0; n /= 10 {
		if n%10 == 0 || num%(n%10) != 0 {
			return false
		}
	}
	return true
}
