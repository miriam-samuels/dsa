package datastructure

import "fmt"

type Node struct {
	Val  interface{}
	Prev *Node
	Next *Node
}

type Queue struct {
	Head *Node
	Tail *Node
	Size int
}

type IQueue interface {
	Enqueue()
	Dequeue()
	Display()
	Exchange()
}

func (q *Queue) Enqueue(val interface{}) {
	newNode := &Node{Val: val}
	if q.Head == nil {
		q.Head = newNode
		q.Tail = newNode
	} else {
		newNode.Prev = q.Tail
		q.Tail.Next = newNode
		q.Tail = newNode
	}

	q.Size++
}

func (q *Queue) Dequeue()*Node {
	if q.Head == nil {
		return nil
	}
	var dequeued = q.Head;
	if q.Head != nil {
		q.Head = dequeued.Next
		q.Head.Prev = nil
		q.Size--
	}
	return dequeued
}

func (q *Queue) Display() {
	current := q.Head
	for {
		fmt.Printf("%v \n", current)

		if current.Next == nil {
			break
		} else {
			current = current.Next
		}
	}
}

func (q *Queue) Exchange(val1, val2 interface{}) {
    if q.Head == nil || q.Head.Next == nil {
        return
    }

    // Find the nodes with the given values
    var node1, node2 *Node
    current := q.Head
    for current != nil {
        if current.Val == val1 {
            node1 = current
        }
        if current.Val == val2 {
            node2 = current
        }
        current = current.Next
    }

    // If either Node is not found, return
    if node1 == nil || node2 == nil {
        return
    }

    // Swap the values of the nodes
    node1.Val, node2.Val = node2.Val, node1.Val
}


// type Node struct {
// 	Val  interface{}
// 	Next *node
// }

// type Queue struct {
// 	Head *node
// 	Tail *node
// 	Size int
// }

// func (q *Queue) enqueue(val interface{}) {
// 	newNode := &node{Val: val}
// 	if q.Head == nil {
// 		q.Head = newNode
// 		q.Tail = newNode
// 	} else {
// 		q.Tail.Next = newNode
// 		q.Tail = newNode
// 	}

// 	q.Size++
// }

// func (q *Queue) dequeue() {
// 	if q.Head != nil {
// 		q.Head = q.Head.Next
// 		q.Size--
// 	}
// }
