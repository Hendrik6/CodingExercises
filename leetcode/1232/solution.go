package leetcode

// 1232. Check If It Is a Straight Line
// You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point.
// Check if these points make a straight line in the XY plane.

func checkStraightLine(coordinates [][]int) bool {
	for i := 2; i < len(coordinates); i++ {
		x1, x2, x3 := coordinates[i-2][0], coordinates[i-1][0], coordinates[i][0]
		y1, y2, y3 := coordinates[i-2][1], coordinates[i-1][1], coordinates[i][1]
		if (y2-y1)*(x3-x2) != (y3-y2)*(x2-x1) {
			return false
		}
	}
	return true
}
