package main

type vec3 struct {
	x, y, z float64
}

func newVec3(value float64) vec3 {
	return vec3{value, value, value}
}

func newVec3XVec2(x float64, v vec2) vec3 {
	return vec3{x, v.x, v.y}
}

func newVec3XYZ(x, y, z float64) vec3 {
	return vec3{x, y, z}
}

func (v vec3) add(other vec3) vec3 {
	return vec3{v.x + other.x, v.y + other.y, v.z + other.z}
}

func (v vec3) subtract(other vec3) vec3 {
	return vec3{v.x - other.x, v.y - other.y, v.z - other.z}
}

func (v vec3) multiply(other vec3) vec3 {
	return vec3{v.x * other.x, v.y * other.y, v.z * other.z}
}

func (v vec3) divide(other vec3) vec3 {
	return vec3{v.x / other.x, v.y / other.y, v.z / other.z}
}

func (v vec3) negate() vec3 {
	return vec3{-v.x, -v.y, -v.z}
}
