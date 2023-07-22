package main

import (
	"fmt"
	"time"
)

const (
	H       = 24
	W       = 80
	gridCap = W*H + 1
)

func main() {
	start(W, H)
}

func start(w, h int) {
	g := newGraph(w, h)
	count := 30
	g.Fill(count)
	endNodeID := g.GetEndAddress()
	for {
		grid := make([]byte, gridCap)

		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				grid[i+j*w] = '*'
			}
		}

		var runner byte = 'R'

		for i, node := range g.Nodes {
			var char byte = 'A'
			if i == 0 {
				char = runner
			}
			grid[node.X+node.Y*w] = char
		}

		if !DFS(grid, g.Nodes[0], endNodeID) {
			fmt.Printf("not found path between %v and %v \n", 1, endNodeID)
		} else {
			fmt.Printf("found path betwwen %v and %v \n", 1, endNodeID)
		}
		endNodeID = g.GetEndAddress()
		time.Sleep(1 * time.Second)
		g.Fill(count)
	}

}
