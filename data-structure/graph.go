package datastructure
// using adjacency list
type GraphNode struct {
	Val interface{}
	neighbors []*GraphNode
}

type Graph struct {
	Head *GraphNode

}

type IGraph interface {
	AddVertice()
	RemoveVertice()
	AddEdge()
	RemoveEdge()
}