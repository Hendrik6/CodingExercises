
pub contract Solution {
    //Based of of multiply this should be changed in order to fix the two sum problem.
    // Public function that multiplies 2 given values
    pub fun twoSum(_ x: [Int],_ y: Int): [Int] {
        var ret: [Int] = []
        var index = 0
        for one in x {
            var innerIndex = 0
            for two in x {
            if one + two == y {
                ret.append(index)
                ret.append(innerIndex)

                return ret
            }
            innerIndex = innerIndex + 1
            } 
            index = index + 1
       }
    return ret
    }
}
 