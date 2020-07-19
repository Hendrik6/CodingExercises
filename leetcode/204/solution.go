package leetcode

//204. Count Primes
//Count the number of prime numbers less than a non-negative number, n.

func countPrimes(n int) (cnt int) {
	isNotPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			cnt++
			for j := i * i; j < n; j += i {
				isNotPrime[j] = true
			}
		}
	}
	return
}
