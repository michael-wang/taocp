// Volume 1, pp xvii, exercise 4.
package main

import (
	"fmt"
	"math"
)

func pows(s, t, n int) map[int]float64 {
	vv := map[int]float64{}
	var v float64
	for i := s; i <= t; i++ {
		v = math.Pow(float64(i), float64(n))
		vv[i] = v
	}
	return vv
}

func main() {
	const max = 100.0
	const n = 2.0
	fmt.Printf("For n=%.f, find x, y, z WHERE x^n + y^n = z^n (max value to seatch: %.f)\n", n, max)
	fmt.Println("Notice: when n > 2, expect NO integer x, y, z can be found to satisfy the equation.")

	xns := pows(1.0, max, n)
	yns := pows(1.0, max, n)
	zns := pows(1.0, max, n)

	for x := 1; x <= max; x++ {
		for y := 1; y <= max; y++ {
			for z := 1; z <= max; z++ {
				if zns[z] == (xns[x] + yns[y]) {
					fmt.Printf("x: %d, y: %d, z: %d\n", x, y, z)
				}
			}
		}
	}
}

