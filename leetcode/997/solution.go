package leetcode

// In a town, there are N people labelled from 1 to N.  There is a rumor that one of these people is secretly the town judge.
// If the town judge exists, then:
// The town judge trusts nobody.
// Everybody (except for the town judge) trusts the town judge.
// There is exactly one person that satisfies properties 1 and 2.
// You are given trust, an array of pairs trust[i] = [a, b] representing that the person labelled a trusts the person labelled b.
// If the town judge exists and can be identified, return the label of the town judge.  Otherwise, return -1.

func findJudge(N int, trust [][]int) int {
	if len(trust) == 0 {
		return 1
	}
	people := make(map[int]int)
	judge := make(map[int]int)
	for _, v := range trust {
		people[v[0]]++
		judge[v[1]]++
	}
	for i, v := range judge {
		if _, ok := people[i]; !ok && v == N-1 {
			return i
		}
	}
	return -1
}
