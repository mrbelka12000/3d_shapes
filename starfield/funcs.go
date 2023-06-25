package main

func ellipse(x1, y1, x2, y2 int) (int, int) {
	var xStep = func(t int) int {
		return x1 + t*(x2-x1)
	}
	var yStep = func(t int) int {
		return y1 + t*(y2-y1)
	}

	return xStep(3), yStep(3)
}

func generateLine(from, to, max int) []int {
	var arr []int
	if from > to {
		for i := from; i >= to; i-- {
			if i <= 0 {
				break
			}
			arr = append(arr, i)
		}
	} else {
		for i := from; i <= to; i++ {
			if i >= max {
				break
			}
			arr = append(arr, i)
		}
	}
	return arr
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
