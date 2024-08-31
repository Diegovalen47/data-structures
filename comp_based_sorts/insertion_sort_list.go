package compbasedsorts

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func RunInsertionSortList() {
	FourthNode := ListNode{
		Val:  3,
		Next: nil,
	}
	ThirdNode := ListNode{
		Val:  1,
		Next: &FourthNode,
	}
	SecondNode := ListNode{
		Val:  2,
		Next: &ThirdNode,
	}
	FirstNode := ListNode{
		Val:  4,
		Next: &SecondNode,
	}
	fmt.Println("input", printLinkedList(&FirstNode))
	fmt.Println("output", printLinkedList(insertionSortList(&FirstNode)))
}

/*
	Given the head of a singly linked list, sort the list using insertion sort, and return the sorted list's head.

The steps of the insertion sort algorithm:

Insertion sort iterates, consuming one input element each repetition and growing a sorted output list.
At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list and inserts it there.
It repeats until no input elements remain.
*/
func insertionSortList(head *ListNode) *ListNode {

	node := head
	var auxList []int
	for node != nil {

		auxList = append(auxList, node.Val)
		for i := 1; i < len(auxList); i++ {
			currentIndex := i

			for currentIndex > 0 && auxList[currentIndex-1] > auxList[currentIndex] {
				auxList[currentIndex-1], auxList[currentIndex] = auxList[currentIndex], auxList[currentIndex-1]
				currentIndex--
			}

		}

		node = node.Next
	}

	return fromSliceToLinkedListHead(auxList)
}

func fromSliceToLinkedListHead(list []int) *ListNode {

	var lastNode *ListNode
	for i := len(list) - 1; i >= 0; i-- {
		if lastNode == nil {
			lastNode = &ListNode{
				Val:  list[i],
				Next: nil,
			}
			continue
		}

		node := ListNode{
			Val:  list[i],
			Next: lastNode,
		}

		lastNode = &node
	}

	return lastNode
}

func printLinkedList(head *ListNode) []int {

	var intList []int
	node := head

	for node != nil {
		intList = append(intList, node.Val)
		node = node.Next
	}
	return intList
}
