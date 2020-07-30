package leetcode

import "sort"

// 4. Median of Two Sorted Arrays
// There are two sorted arrays nums1 and nums2 of size m and n respectively.
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
// You may assume nums1 and nums2 cannot be both empty.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)

	sort.Ints(nums1)

	if len(nums1)%2 != 0 {
		return float64(nums1[len(nums1)/2])
	}

	float1, float2 := (float64(nums1[(len(nums1)/2)-1])), (float64(nums1[(len(nums1) / 2)]))

	return (float1 + float2) / 2
}
