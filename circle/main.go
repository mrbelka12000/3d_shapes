package main

import (
	"fmt"
	"math"
)

const (
	H = 24
	W = 80
)

func main() {
	draw(H, W)
}

func draw(h, w int) {
	gradient := []byte{' ', '.', 'a', '@'}
	gradientSize := len(gradient) - 1
	var t float64
	for {
		aspect := float64(w) / float64(h)
		pixelAspect := 11.0 / 24.0
		shape := make([]byte, w*h+1)
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				//x := float64(i)/float64(w)*2.0 - 1.0
				//y := float64(j)/float64(h)*2.0 - 1.0
				uv := vec2{float64(i), float64(j)}.divide(vec2{float64(w), float64(h)}).multiply(vec2{2.0, 2.0}).subtract(vec2{1.0, 1.0})
				uv.x *= aspect * pixelAspect
				uv.x += math.Sin(t * 0.0001)
				//ro := vec3{-5, 0, 0}
				//rd := newVec3XVec2(0, uv)
				dist := math.Sqrt(uv.x*uv.x + uv.y*uv.y)
				color := int(1.0 / dist)
				color = clVal(color, 0, gradientSize)
				shape[i+j*w] = gradient[color]
			}
		}
		t++
		fmt.Println(string(shape))
	}
}

func clVal(val, min, max int) int {
	if val < min {
		return min
	} else if val >= max {
		return max
	}
	return val
}
