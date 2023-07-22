package main

import "math/rand"

func getRandomCoordinates(maxW, maxH int) coordinates {
	return coordinates{
		X: rand.Intn(maxW),
		Y: rand.Intn(maxH),
	}
}

func cleanGrid(grid []byte, coords []int) {
	for _, coord := range coords {
		grid[coord] = '*'
	}
}
