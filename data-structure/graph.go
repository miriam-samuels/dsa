package datastructure

//E <= V^2
// using adjacency list
// directed

type Neighbors []int

// type Graph struct{
// 	Val int
// 	neighbors []*Graph
// }

// using hash map
type Graph map[int]Neighbors

type IGraph interface {
	AddVertex()
	RemoveVertex()
	AddEdge()
	RemoveEdge()
}

func (g *Graph) AddVertex(vertex int) {
	if _, ok := (*g)[vertex]; !ok {
		(*g)[vertex] = make([]int, 0)
	}
}

func (g *Graph) RemoveVertex(vertex int) {
	delete((*g), vertex)
	for _, neighbors := range *g {
		for i, neighbor := range neighbors {
			if neighbor == vertex {
				(*g)[neighbor] = append(neighbors[:i], neighbors[i+1:]...)
				return
			}
		}
	}
}

func (g *Graph) AddEdge(src, dst int) {
	if _, ok := (*g)[src]; ok {
		(*g)[src] = append((*g)[src], dst)
	}
}

func (g *Graph) RemoveEdge(src, dst int) {
	if _, ok := (*g)[src]; ok {
		for i, neighbor := range (*g)[src] {
			if neighbor == dst {
				(*g)[src] = append((*g)[src][:i], (*g)[src][i+1:]...)
			}
		}
	}
}

func (m *Graph) DFS(startNode, targetNode int, visited [][]bool) {

}

func (m *Graph) BFS(startRow, startCol int) {

}
