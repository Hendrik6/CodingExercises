import Solution from 0x01cf0e2f2f715450
transaction {
	prepare(acct: AuthAccount) {
    }
	
	execute {
        var numbers1 = [11, 7, 13, 2]
        var expected1 = [1,3]
        var result1: [Int] = []
	    result1 = Solution.twoSum(numbers1, 9)

        log("First output expected / result")
        log(expected1)
	  	log(result1)

        ///////////

        var numbers2 = [3, 8, 13, 24]
        var expected2 = [0,2]
        var result2: [Int] = []
        result2 = Solution.twoSum(numbers2, 16)

        log("Second output expected / result")
        log(expected2)
	  	log(result2)
	}
}
 