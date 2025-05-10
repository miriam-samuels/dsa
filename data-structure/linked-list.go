package datastructure

import "fmt"

type ListNode struct {
	Val  interface{}
	Prev *ListNode
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

type ILinkedList interface {
	Push()
	Pop()
	Insert()
	Remove()
	Len()
	Update()
	Display()
	Get()
}

func (l *LinkedList) Push(val interface{}) { //add to end
	newNode := &ListNode{Val: val}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Prev = l.Tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

func (l *LinkedList) Pop() { // remove last element
	if l.Head != nil {
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
	}
}

func (l *LinkedList) Insert(val interface{}, pos int) {
	newNode := &ListNode{Val: val}

	current := l.Head
	for i := 1; i < pos; i++ {
		current = current.Next
	}

	newNode.Next = current
	newNode.Prev = current.Prev
	current.Prev.Next = newNode
	current.Prev = newNode

}

func (l *LinkedList) Remove(pos int) {
	current := l.Head
	for i := 1; i < pos; i++ {
		current = current.Next
	}
	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev

}

func (l *LinkedList) Update(val interface{}, pos int) {
	current := l.Head
	for i := 1; i < pos; i++ {
		current = current.Next
	}

	current.Val = val
}

func (l *LinkedList) Find() {

}

func (l *LinkedList) Sort() {

}

func (l *LinkedList) Len() int {
	var count int = 1
	current := l.Head
	for {
		if current.Next == nil {
			break
		} else {
			current = current.Next
		}
		count++
	}
	return count
}

func (l *LinkedList) Display() {
	current := l.Head
	for {
		fmt.Printf("%v \n", current)

		if current.Next == nil {
			break
		} else {
			current = current.Next
		}
	}
}
