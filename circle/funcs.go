package main

import "math"

func clamp(value, min, max float64) int {
	return int(math.Max(math.Min(value, max), min))
}

func sign(a float64) float64 {
	var i, j float64
	if a < 0 {
		i = 1
	}
	if 0 < a {
		j = 1
	}
	return j - i
}

func step(edge, x float64) float64 {
	if x > edge {
		return 1.0
	}
	return 0.0
}

func lengthV2(v vec2) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func lengthV3(v vec3) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func norm(v vec3) vec3 {
	return v.divide(newVec3(lengthV3(v)))
}

func dot(a, b vec3) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func abs(v vec3) vec3 {
	return vec3{math.Abs(v.x), math.Abs(v.y), math.Abs(v.z)}
}

func signVec(v vec3) vec3 {
	return vec3{sign(v.x), sign(v.y), sign(v.z)}
}

func stepVec(edge, v vec3) vec3 {
	return vec3{step(edge.x, v.x), step(edge.y, v.y), step(edge.z, v.z)}
}

func reflect(rd, n vec3) vec3 {
	return rd.subtract(n.multiply(newVec3(2 * dot(n, rd))))
}

func rotateX(a vec3, angle float64) vec3 {
	b := a
	b.z = a.z*math.Cos(angle) - a.y*math.Sin(angle)
	b.y = a.z*math.Sin(angle) + a.y*math.Cos(angle)
	return b
}

func rotateY(a vec3, angle float64) vec3 {
	b := a
	b.x = a.x*math.Cos(angle) - a.z*math.Sin(angle)
	b.z = a.x*math.Sin(angle) + a.z*math.Cos(angle)
	return b
}

func rotateZ(a vec3, angle float64) vec3 {
	b := a
	b.x = a.x*math.Cos(angle) - a.y*math.Sin(angle)
	b.y = a.x*math.Sin(angle) + a.y*math.Cos(angle)
	return b
}

func sphere(ro, rd vec3, r float64) vec2 {
	b := dot(ro, rd)
	c := dot(ro, ro) - r*r
	h := b*b - c
	if h < 0.0 {
		return vec2{-1.0, -1.0}
	}
	h = math.Sqrt(h)
	return vec2{-b - h, -b + h}
}

func box(ro, rd, boxSize vec3) (vec2, vec3) {
	m := vec3{1.0, 1.0, 1.0}.divide(rd)
	n := m.multiply(ro)
	k := abs(m).multiply(boxSize)
	t1 := n.negate().subtract(k)
	t2 := n.negate().add(k)
	tN := math.Max(math.Max(t1.x, t1.y), t1.z)
	tF := math.Min(math.Min(t2.x, t2.y), t2.z)
	if tN > tF || tF < 0.0 {
		return vec2{-1.0, -1.0}, vec3{}
	}
	yzx := vec3{t1.y, t1.z, t1.x}
	zxy := vec3{t1.z, t1.x, t1.y}
	outNormal := signVec(rd).multiply(stepVec(yzx, t1)).multiply(stepVec(zxy, t1))
	return vec2{tN, tF}, outNormal
}

func plane(ro, rd, p vec3, w float64) float64 {
	return -(dot(ro, p) + w) / dot(rd, p)
}
