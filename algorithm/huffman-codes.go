package algorithm
const (
	s = "AAAAABBBCCCCCAAAAADDDDD111118888833778782222"
)
type Node struct{
	Char rune
	Freq int
	Left,Right *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len(){
	return len(pq)
}

func (pq PriorityQueue) IsLess(i,j int)bool{
	return pq[i].Freq < pq[j].Freq
}

func (pq PriorityQueue) Swap(i,j int){
	pq[i],pq[j] = pq[j],pq[i]
}

func (pq *PriorityQueue) Push(x interface{}){
	*pq =  append(*pq, x.(*Node))
}
charFrequency := map[rune]int

for _, c := range s {
	charFrequency[c] += 1 
}

