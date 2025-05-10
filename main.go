package main

import (
	"fmt"

	structure "github.com/miriam-samuels/go-dsa/data-structure"
)

func main() {
	fmt.Println((2-1) / 2)
	b := &structure.BinaryTree{}
	b.Insert(0)
	b.Insert(10)
	b.Insert(11)
	b.Insert(8)
	b.Insert(1)
	b.Insert(6)
	b.Insert(9)

	b.Search(9)
	// b.InOrder()
	// b.PreOrder()
	// b.PostOrder()

	// q := &structure.Queue{}
	// q.Enqueue(1)
	// q.Enqueue(2)
	// q.Enqueue(3)
	// q.Enqueue(4)
	// q.Dequeue()
	// q.Display()

	// ll := &structure.LinkedList{}

	// ll.Push(1)
	// ll.Push(2)
	// ll.Push(3)
	// ll.Pop()
	// ll.Push(3)
	// ll.Push(4)
	// ll.Remove(2)
	// ll.Insert("22", 2)
	// fmt.Printf("%d \n", ll.Len())
	// ll.Display()
}
