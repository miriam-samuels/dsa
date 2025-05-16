package datastructure

import "fmt"

// var matrix = &Matrix{
// 	{1, 1, 0, 0, 0},
// 	{1, 1, 0, 1, 1},
// 	{0, 0, 0, 1, 1},
// 	{0, 0, 0, 0, 0},
// }

type Matrix [][]int

type IMatrix interface {
	DFS()
	BFS()
	NumIslands()
}

// Q1: Count the unique paths from the top left to the bottom right. A single path may only move along 0's and can't visit the same cell more than once.
// Recursion uses the call stack, which behaves like a stack (LIFO).(explores farthest then backtracks)
func (m *Matrix) DFS(currRow, currCol int, visited [][]bool) int { 
	rows, cols := len(*m), len((*m)[0])

	//check if out of bounds 0,0 starting point
	if (currRow < 0 || currRow >= rows) || (currCol < 0 || currCol >= cols) {
		return 0
	}

	// check if  row/col has been visited
	if visited[currRow][currCol] {
		return 0
	}

	// check if it's an available path
	if (*m)[currRow][currCol] == 1 {
		return 0
	}

	// mark as visited
	visited[currRow][currCol] = true

	// for q1 we check if we successfully got to the end (bottom right)
	if (currRow == rows-1) && (currCol == cols-1) {
		return 1 // count as a unique complete path
	}

	uniqueCount := 0

	// explore all other paths
	uniqueCount += m.DFS(currRow-1, currCol, visited) // up
	uniqueCount += m.DFS(currRow+1, currCol, visited) // down
	uniqueCount += m.DFS(currRow, currCol-1, visited) // left
	uniqueCount += m.DFS(currRow, currCol+1, visited) // right

	return uniqueCount
}

// uses a queue FIFO (Explores closest before farthese)
func (m *Matrix) BFS(startRow, startCol int) {
	// ensure matrix lenght is not 0
	if len(*m) == 0 {
		return 
	}

	// number of rows and cols
	rows, cols := len(*m), len((*m)[0])

	// track visited coordinates
	visited := make([][]bool, rows)
	// map visited entry for each row and col
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// use a queue because it's FIFO
	queue := &Queue{}

	// add starting coordinates or start coordinates 0,0 to array of len 2
	queue.Enqueue([2]int{startRow, startCol})

	// mark start as visited
	visited[startRow][startCol] = true

	// as long as the queue is not empty
	for queue.Size > 0 {
		// remove from queue
		node := queue.Dequeue()
		if node == nil {
			break
		}

		// store val of node as cell and assert type array of len 2
		cell := node.Val.([2]int)
		currRow, currCol := cell[0], cell[1]
		fmt.Printf("Visiting cell (%d, %d) with value %d\n", currRow, currCol, (*m)[currRow][currCol])

		// move in all directions r,c
		directions := [4][2]int{
			{0, 1},  //right
			{-1, 0}, //up
			{0, -1}, // left
			{1, 0},  //down
		}

		// loop through for all directions
		for _, direction := range directions {
			// get location of neighbor at direction
			neighborRow, neighborCol := currRow+direction[0], currCol+direction[1]

			// check if neighbor is out of matrix bounds
			if (neighborRow < 0 || neighborRow >= rows) || (neighborCol < 0 || neighborCol >= cols) {
				continue
			}

			// check if neighbor has been visited
			if visited[neighborRow][neighborCol] {
				continue
			}

			queue.Enqueue([2]int{neighborRow, neighborRow})
			visited[neighborRow][neighborCol] = true
		}
	}

}


func (matrix *Matrix) NumIslands() int {
	// var matrix = &Matrix{
	// 	{1, 1, 0, 0, 0},
	// 	{1, 1, 0, 1, 1},
	// 	{0, 0, 0, 1, 1},
	// 	{0, 0, 0, 0, 0},
	// }

	// ensure matrix lenght is not 0
	if len(*matrix) == 0 {
		return 0
	}

	rows, cols := len(*matrix), len((*matrix)[0])
	// map visited entry for each row and col
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// count islands
	islands := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (*matrix)[i][j] == 1 && !visited[i][j] {
				matrix.DFS(i, j, visited)
				islands++
			}
		}
	}
	return islands
}
