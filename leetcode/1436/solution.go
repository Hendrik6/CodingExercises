package leetcode

// 1436. Destination City
// You are given the array paths, where paths[i] = [cityAi, cityBi] means there exists a direct path going from cityAi to cityBi.
// Return the destination city, that is, the city without any path outgoing to another city.
// It is guaranteed that the graph of paths forms a line without any loop, therefore, there will be exactly one destination city.

func destCity(paths [][]string) string {
	tmp := make(map[string]int)
	for _, v := range paths {
		for i := 0; i < 2; i++ {
			tmp[v[i]]++
		}
	}
	for _, v := range paths {
		for i := 0; i < 2; i++ {
			if tmp[v[i]] == 1 && i == 1 {
				return v[i]
			}
		}
	}
	return ""
}
