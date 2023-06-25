package main

import (
	"math/rand"
)

type Star struct {
	X, Y, Z int
}

func NewStar(maxW, maxY int) *Star {
	return &Star{
		X: rand.Intn(maxW),
		Y: rand.Intn(maxY),
		Z: maxW,
	}
}

func (s *Star) Update() {
	s.X = rand.Intn(W)
	s.Y = rand.Intn(H)
	s.Z = W
}

func (s *Star) Show(cx, cy int, grid []byte) {
	xStep, yStep := ellipse(cx, cy, s.X, s.Y)
	grid[s.X+s.Y*W] = '*'
	gridCap := W*H + 1

	arrx := generateLine(s.X, xStep, W)
	arry := generateLine(s.Y, yStep, H)
	if len(arrx) == 1 {
		if arrx[0] != W-1 && arrx[0] != 1 {
			for i := 0; i < len(arry); i++ {
				arrx = append(arrx, arrx[0])
			}
		}
	}
	if len(arry) == 1 {
		if arry[0] != H-1 && arry[0] != 1 {
			for i := 0; i < len(arrx)/2; i++ {
				arry = append(arry, arry[0])
			}
		}
	}
	ln := max(len(arrx), len(arry))
	for i := 0; i < ln; i++ {
		if i >= len(arrx) || i >= len(arry) {
			break
		}
		x, y := arrx[i], arry[i]
		pos := x + y*W
		if pos >= gridCap || pos < 0 {
			break
		}

		if grid[pos] != '*' {
			var char byte = 'o'
			grid[pos] = char
		}
	}

	if xStep > s.X {
		s.X += speed
	} else if xStep < s.X {
		s.X -= speed
	}
	if yStep > s.Y {
		s.Y += speed
	} else if yStep < s.Y {
		s.Y -= speed
	}
	pos := s.X + s.Y*W
	if pos >= gridCap || pos < 0 {
		s.Rand()
	}
}

func (s *Star) Rand() {
	s.Update()
}
