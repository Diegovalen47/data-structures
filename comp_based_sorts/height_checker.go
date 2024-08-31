// Bubble sort
package compbasedsorts

import "fmt"

func RunHeightChecker() {
	testCase := []int{1, 1, 4, 2, 1, 3}
	fmt.Println("input", testCase)
	fmt.Println("output", heightChecker(testCase))
}

/*
	A school is trying to take an annual photo of all the students. The students are asked to stand in a single file line in non-decreasing order by height. Let this ordering be represented by the integer array expected where expected[i] is the expected height of the ith student in line.

You are given an integer array heights representing the current order that the students are standing in. Each heights[i] is the height of the ith student in line (0-indexed).

Return the number of indices where heights[i] != expected[i].
*/
func heightChecker(heights []int) int {
	lenght := len(heights)
	orderedHeights := make([]int, lenght)
	copy(orderedHeights, heights)

	hasSwapped := true
	for hasSwapped {
		hasSwapped = false
		for i := 0; i < lenght-1; i++ {
			if orderedHeights[i] > orderedHeights[i+1] {
				orderedHeights[i+1], orderedHeights[i] = orderedHeights[i], orderedHeights[i+1]
				hasSwapped = true
			}
		}
	}

	different := 0
	for j := 0; j < lenght; j++ {
		if heights[j] != orderedHeights[j] {
			different++
		}
	}

	return different
}
