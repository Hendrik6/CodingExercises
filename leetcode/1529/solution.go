package leetcode

// 1529. Bulb Switcher IV

// There is a room with n bulbs, numbered from 0 to n-1, arranged in a row from left to right. Initially all the bulbs are turned off.
// Your task is to obtain the configuration represented by target where target[i] is '1' if the i-th bulb is turned on and is '0' if it is turned off.
// You have a switch to flip the state of the bulb, a flip operation is defined as follows:
// Choose any bulb (index i) of your current configuration.
// Flip each bulb from index i to n-1.
// When any bulb is flipped it means that if it is 0 it changes to 1 and if it is 1 it changes to 0.
// Return the minimum number of flips required to form target.

func minFlips(target string) int {
	simplified := []byte{target[0]}
	last := target[0]
	for i := 1; i < len(target); i++ {
		if target[i] == last {
			continue
		}
		last = target[i]
		simplified = append(simplified, target[i])
	}
	if simplified[0] == '0' {
		return len(simplified) - 1
	}
	return len(simplified)
}
