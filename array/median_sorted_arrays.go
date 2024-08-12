package array

import (
	"math"
)

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	m := len(nums1)
	n := len(nums2)
	sum := m + n

	mergedArray := make([]int, 0, sum)

	i := 0
	j := 0
	for k := 0; k < cap(mergedArray); k++ {
		if i == m {
			mergedArray = append(mergedArray, nums2[j:]...)
			break
		} else if j == n {
			mergedArray = append(mergedArray, nums1[i:]...)
			break
		}

		if nums1[i] <= nums2[j] {
			mergedArray = append(mergedArray, nums1[i])
			i++
		} else {
			mergedArray = append(mergedArray, nums2[j])
			j++
		}
	}

	if sum%2 == 0 {
		leftMedian := mergedArray[int(math.Floor(float64(sum/2)))-1]
		rightMedian := mergedArray[int(math.Ceil(float64(sum/2)))]

		return float64(float64(leftMedian+rightMedian) / 2)
	}

	return float64(mergedArray[(sum / 2)])
}
