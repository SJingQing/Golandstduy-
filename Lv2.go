package main

import (
	"fmt"
	"math"
)

func mathoperation(x float64, y float64) float64 {
	return x * math.Pow(y, 2)
}
func main() {
	s := mathoperation(math.Pi, 3)
	fmt.Println("圆的面积s=", s)
}
