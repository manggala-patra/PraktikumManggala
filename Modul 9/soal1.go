package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y float64
}

type Lingkaran struct {
	pusat Titik
	r     float64
}

func jarak(a, b Titik) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func diDalam(l Lingkaran, p Titik) bool {
	return jarak(l.pusat, p) <= l.r
}

func main() {
	var x1, y1, r1 float64
	var x2, y2, r2 float64
	var px, py float64

	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&px, &py)

	ling1 := Lingkaran{Titik{x1, y1}, r1}
	ling2 := Lingkaran{Titik{x2, y2}, r2}
	titik := Titik{px, py}

	dalam1 := diDalam(ling1, titik)
	dalam2 := diDalam(ling2, titik)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}