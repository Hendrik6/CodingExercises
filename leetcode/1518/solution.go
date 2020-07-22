package leetcode

// 1518. Water Bottles

// Given numBottles full water bottles, you can exchange numExchange empty water bottles for one full water bottle.
// The operation of drinking a full water bottle turns it into an empty bottle.
// Return the maximum number of water bottles you can drink.

func numWaterBottles(numBottles int, numExchange int) int {
	sum, emptyBottles := 0, 0
	for numBottles > 0 || emptyBottles >= numExchange {
		sum += numBottles
		emptyBottles += numBottles

		numBottles = emptyBottles / numExchange
		emptyBottles = emptyBottles % numExchange
	}
	return sum
}
