package datastructure

//priority min queue / heap
// 1 index
// left child  =  i * 2
// right child  = i * 2 + 1
// parent = math.floor(i / 2)

// 0 index
// left child = i * 2 + 1
// right child  = i * 2 + 2
// parent = math.floor(i - 1 / 2)

type MinHeap []int

type IHeap interface {
	Push()
	Pop()
	Swap()
	heapifyUp()
	heapifyDown()
}

func (h *MinHeap) Push(val int) {
	*h = append(*h, val)
	h.heapifyUp()
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) heapifyUp() {
	i := len(*h) - 1

	for i > 0 {
		parent := (i - 1) / 2

		if (*h)[parent] > (*h)[i] {
			h.Swap(parent, i)
			i = parent
		}
	}
}

func (h *MinHeap) heapifyDown() {
	size := len(*h) - 1
	parent := 0
	for parent < size {
		leftChild := parent*2 + 1
		rightChild := parent*2 + 2

		if parent > leftChild {
			h.Swap(parent, leftChild)
			parent = leftChild
		}
		if parent > rightChild {
			h.Swap(parent, rightChild)
			parent = rightChild
		}
	}
}

func (q *Queue) Push(val int) {
	q.Enqueue(val)
}

func (q *Queue) Swap(val1, val2 int) {
	q.Exchange(val1, val2)
}
