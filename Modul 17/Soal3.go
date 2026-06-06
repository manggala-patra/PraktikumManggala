package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int

	fmt.Scan(&n)

	rand.Seed(time.Now().UnixNano())

	var A, B, C, D int

	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if x < 0.5 && y < 0.5 {
			A++
		} else if x >= 0.5 && y < 0.5 {
			B++
		} else if x >= 0.5 && y >= 0.5 {
			C++
		} else {
			D++
		}
	}

	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", float64(A)*0.0001)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", float64(B)*0.0001)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", float64(C)*0.0001)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", float64(D)*0.0001)
}