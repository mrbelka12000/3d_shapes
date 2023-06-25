package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	H        = 24
	W        = 80
	capacity = 100
)

func main() {
	rand.Seed(time.Now().UnixNano())
	run(W, H)
}

func run(w, h int) {
	stars := make([]*Star, capacity)
	for i := 0; i < capacity; i++ {
		stars[i] = NewStar(w, h)
	}
	cx, cy := (w-1)/2, (h-1)/2
	gridCap := w*h + 1
	for {
		grid := make([]byte, gridCap)

		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				grid[i+j*w] = '.'
			}
		}

		for i := 0; i < capacity; i++ {
			stars[i].Show(cx, cy, grid)
		}
		grid[cx+cy*w] = '0'
		fmt.Println(string(grid))
		time.Sleep(500 * time.Millisecond)
	}
}
