// Selection Sort
package compbasedsorts

import "fmt"

func RunSortColors() {
	testCase := []int{2, 0, 2, 1, 1, 0}
	fmt.Println("input", testCase)
	fmt.Println("output", sortColors(testCase))
}

func sortColors(nums []int) []int {

	newNumbers := make([]int, len(nums))
	copy(newNumbers, nums)

	for i := 0; i < len(newNumbers); i++ {
		minIndex := i
		for j := i + 1; j < len(newNumbers); j++ {
			if newNumbers[j] < newNumbers[minIndex] {
				minIndex = j
			}
		}

		temp := newNumbers[minIndex]
		newNumbers[minIndex] = newNumbers[i]
		newNumbers[i] = temp
	}
	return newNumbers
}
