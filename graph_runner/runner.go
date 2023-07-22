package main

func DrawRunner(grid []byte, x, y int) (coords []int) {
	for i := x - 1; i <= x+1 && i >= 0 && y < W; i++ {
		pos := i + y*W
		if pos > gridCap {
			continue
		}
		coords = append(coords, pos)
		grid[pos] = '@'
	}
	for i := y - 1; i <= y+1 && i >= 0 && i < H; i++ {
		pos := x + i*W
		if pos > gridCap {
			continue
		}
		coords = append(coords, pos)
		grid[pos] = '@'
	}
	return coords
}
