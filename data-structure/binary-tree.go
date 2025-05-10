package datastructure

import "fmt"

type BinaryNode struct {
	Val   int
	Left  *BinaryNode
	Right *BinaryNode
}

type BinaryTree struct {
	Root *BinaryNode
	Size int
}

type ITree interface {
	Insert()
	Search()
	Delete()
	PreOrder() // depth first
	InOrder() // depth first
	PostOrder() // depth first
	LevelOrder() // breadth first
}

func (b *BinaryTree) Insert(val int) {
	if b.Root == nil {
		b.Root = &BinaryNode{}
		b.Root.Val = val
	} else {
		var assignNode func(val int, node *BinaryNode)
		assignNode = func(val int, node *BinaryNode) {
			if val < node.Val {
				if node.Left == nil {
					node.Left = &BinaryNode{}
					node.Left.Val = val
				} else {
					assignNode(val, node.Left)
				}
			} else {
				if node.Right == nil {
					node.Right = &BinaryNode{}
					node.Right.Val = val
				} else {
					assignNode(val, node.Right)
				}
			}
		}
		assignNode(val, b.Root)
	}

	b.Size++
}

func (b *BinaryTree) Search(val int) {
	var checkNode func(node *BinaryNode)
	checkNode = func(node *BinaryNode) {
		if node == nil {
			return
		}
		if val == node.Val {
			fmt.Println("Found", node.Val)
			return
		} else {
			if val < node.Val {
				checkNode(node.Left)
			} else {
				checkNode(node.Right)
			}
		}
	}
	checkNode(b.Root)
}

//TODO implement delete
func (b *BinaryTree) Delete(val int)  {
	var checkNode func(node *BinaryNode) *BinaryNode
	checkNode = func(node *BinaryNode) *BinaryNode {
		if node == nil {
			return nil
		}
		if val == node.Val {
			if node.Left == nil {
				return node.Right
			}

			if node.Right == nil {
				return node.Left
			}
			
		} else {
			if val < node.Val {
				checkNode(node.Left)
			} else {
				checkNode(node.Right)
			}
		}
		return node
	}
	checkNode(b.Root)
}

func (b *BinaryTree) InOrder() {
	fmt.Println("Depth First In Order -- left , root, right")
	var checkNode func(node *BinaryNode)
	checkNode = func(node *BinaryNode) {
		if node == nil {
			return
		}
		checkNode(node.Left)
		fmt.Println(node.Val)
		checkNode(node.Right)
	}

	checkNode(b.Root)
}

func (b *BinaryTree) PreOrder() {
	fmt.Println("Depth First Pre Order -- root, left , right")
	var checkNode func(node *BinaryNode)
	checkNode = func(node *BinaryNode) {
		if node == nil {
			return
		}
		fmt.Println(node.Val)
		checkNode(node.Left)
		checkNode(node.Right)
	}

	checkNode(b.Root)
}

func (b *BinaryTree) PostOrder() {
	fmt.Println("Depth First Post Order -- left , right, root")
	var checkNode func(node *BinaryNode)
	checkNode = func(node *BinaryNode) {
		if node == nil {
			return
		}
		checkNode(node.Left)
		checkNode(node.Right)
		fmt.Println(node.Val)
	}

	checkNode(b.Root)
}

//TODO implement level order traversal
func (b *BinaryTree)LevelOrder()  {
	fmt.Println("Breath First Level Order -- left to right")

	if b.Root == nil {
		return
	}

	queue := &Queue{}

	queue.Enqueue(b.Root.Val) // first elenent

	level := 0
	for queue.Size > 0 {
		level++
		fmt.Println("Level", level)
		for i := 0; i < queue.Size; i++ {
			node := queue.Dequeue()
			fmt.Println(node.Val)
			if node.Prev != nil {
				queue.Enqueue(node.Prev.Val)
			}
			if node.Next != nil {
				queue.Enqueue(node.Next.Val)
			}
		}
	}
	fmt.Println("End of level order traversal")
	fmt.Println("Size of queue", queue.Size)
}
