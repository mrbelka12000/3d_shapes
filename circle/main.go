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
	gradient := []byte{' ', ':', '!', 'r', '(', 'l', '1', 'Z', '4', 'H', '9', 'W', '8', '$', '@'}
	gradientSize := len(gradient) - 1
	var t float64
	for {
		light := norm(newVec3XYZ(math.Sin(t*0.01), math.Cos(t*0.01), -1.0))
		aspect := float64(w) / float64(h)
		pixelAspect := 11.0 / 24.0
		shape := make([]byte, w*h+1)
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				//x := float64(i)/float64(w)*2.0 - 1.0
				//y := float64(j)/float64(h)*2.0 - 1.0
				uv := vec2{float64(i), float64(j)}.divide(vec2{float64(w), float64(h)}).multiply(vec2{2.0, 2.0}).subtract(vec2{1.0, 1.0})
				uv.x *= aspect * pixelAspect
				//uv.x += math.Sin(t * 0.0001)
				ro := vec3{-2, 0, 0}
				rd := norm(newVec3XVec2(1, uv))
				var color int
				intersections := sphere(ro, rd, 1)
				if intersections.x > 0 {
					itPoint := ro.add(rd).multiply(newVec3(intersections.x))
					n := norm(itPoint)
					diff := dot(n, light)
					fmt.Println(n)
					fmt.Println(light)
					fmt.Println()
					color = int(diff) * 14
				}
				//diff := dot(n, light)
				//color = int(diff * 20)
				color = clVal(color, 0, gradientSize)
				//dist := math.Sqrt(uv.x*uv.x + uv.y*uv.y)
				//color := int(1.0 / dist)
				//color = clVal(color, 0, gradientSize)
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
