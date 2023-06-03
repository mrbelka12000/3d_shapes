package main

type vec2 struct {
	x float64
	y float64
}

func newVec2(x float64) vec2 {
	return vec2{x, x}
}

func newVec2XY(x, y float64) vec2 {
	return vec2{x, y}
}

func (v vec2) add(other vec2) vec2 {
	return vec2{v.x + other.x, v.y + other.y}
}

func (v vec2) subtract(other vec2) vec2 {
	return vec2{v.x - other.x, v.y - other.y}
}

func (v vec2) multiply(other vec2) vec2 {
	return vec2{v.x * other.x, v.y * other.y}
}

func (v vec2) divide(other vec2) vec2 {
	return vec2{v.x / other.x, v.y / other.y}
}
