package main

import (
	"fmt"
	"testing"
)

func TestEllipse(t *testing.T) {
	nextX, nextY := ellipse(60, 60, 40, 80)
	fmt.Println(nextX, nextY)
}

func TestGenerateLine(t *testing.T) {
	arr := generateLine(11, -17, W)
	fmt.Println(arr)
}
