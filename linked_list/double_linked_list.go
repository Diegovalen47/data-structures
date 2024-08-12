package linkedlist

import "fmt"

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

type Person struct {
	id   int
	name string
	age  int
	gpa  float64
}

func (dll *DoubleLinkedList) AddNode(data interface{}) {
	newNode := &Node{
		data: data,
		prev: nil,
		next: nil,
	}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

func (dll *DoubleLinkedList) PrintForward() {
	currentNode := dll.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (dll *DoubleLinkedList) PrintReverse() {
	currentNode := dll.tail
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.prev
	}
	fmt.Println()
}

func (dll *DoubleLinkedList) Peek() {
	fmt.Println(dll.head)
}

func MainExecution() {
	dll := DoubleLinkedList{}
	dll.AddNode(Person{id: 1007290916, name: "Valentin Osorio", age: 24, gpa: 4.6})
	dll.AddNode(Person{id: 1007243312, name: "Ana Torres", age: 23, gpa: 4.5})
	dll.AddNode(Person{id: 43469098, name: "Nancy Marin", age: 58, gpa: 4.0})
	dll.AddNode(Person{id: 70901147, name: "Alberto Osorio", age: 63, gpa: 4.03})

	fmt.Println("Doubly Linked List (forward):")
	dll.PrintForward()

	fmt.Println("Doubly Linked List (reverse):")
	dll.PrintReverse()

	fmt.Println("Peek")
	dll.Peek()
}
