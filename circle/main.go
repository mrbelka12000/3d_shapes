package main

import (
	"fmt"
	"math"
	"time"
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
	aspect := float64(w) / float64(h)
	pixelAspect := 11.0 / 24.0
	for {
		light := norm(newVec3XYZ(-0.5, 0.5, -1.0))
		spherePos := newVec3XYZ(0, 3, 0)
		shape := make([]byte, w*h+1)
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				uv := vec2{float64(i), float64(j)}.divide(vec2{float64(w), float64(h)}).multiply(vec2{2.0, 2.0}).subtract(vec2{1.0, 1.0})
				uv.x *= aspect * pixelAspect
				ro := vec3{-6, 0, 0}
				rd := norm(newVec3XVec2(1, uv))
				ro = rotateY(ro, 0.25)
				rd = rotateY(rd, 0.25)
				ro = rotateZ(ro, t*0.01)
				rd = rotateZ(rd, t*0.01)

				diff := 1.0
				for i := 0; i < 5; i++ {
					minIt := 99999.0
					intersections := sphere(ro.subtract(spherePos), rd, 1)
					n := newVec3(0)
					albedo := 1.0
					if intersections.x > 0 {
						itPoint := ro.subtract(spherePos).add(rd).multiply(newVec3(intersections.x))
						// itPoint := rd.multiply(newVec3(intersections.x)).add(rd).subtract(spherePos)
						minIt = intersections.x
						n = norm(itPoint)
					}
					intersections, boxN := box(ro, rd, newVec3(1.5))
					if intersections.x > 0 && intersections.x < minIt {
						minIt = intersections.x
						n = boxN
					}
					// intersections = newVec2(plane(ro, rd, newVec3XYZ(0, 0, -1), 1))
					// if intersections.x > 0 && intersections.x < minIt {
					// 	minIt = intersections.x
					// 	n = newVec3XYZ(0, 0, -1)
					// 	albedo = 0.5
					// }
					if minIt < 99999 {
						diff *= (dot(n, light)*0.5 + 0.5) * albedo
						ro = rd.multiply(newVec3(minIt - 0.01)).add(ro)
						rd = reflect(rd, n)
					} else {
						break
					}

				}
				color := int(math.Round(diff)) * gradientSize
				color = clVal(color, 0, gradientSize)
				shape[i+j*w] = gradient[color]
			}
		}
		time.Sleep(7 * time.Millisecond)
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
