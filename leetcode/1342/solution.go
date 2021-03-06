package leetcode

//1342. Number of Steps to Reduce a Number to Zero
//Given a non-negative integer num, return the number of steps to reduce it to zero.
//If the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.

func numberOfSteps(num int) int {
	count := 0
	for num > 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num--
		}
		count++
	}
	return count
}
