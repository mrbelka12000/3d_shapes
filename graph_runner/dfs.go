package main

import (
	"fmt"
	"time"
)

func DFS(grid []byte, node *Node, endID int) bool {
	time.Sleep(800 * time.Millisecond)

	coords := DrawRunner(grid, node.X, node.Y)

	fmt.Println(string(grid))
	cleanGrid(grid, coords)

	grid[node.X+node.Y*W] = 'A'
	node.Visited = true
	if node.ID == endID {
		return true
	}

	for _, v := range node.Adj {
		if !v.Visited {
			val := DFS(grid, v, endID)
			if val {
				return true
			}
		}
	}
	return false
}
