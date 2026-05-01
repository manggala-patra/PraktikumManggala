package main

import "fmt"

const NMAX = 1000

func main() {
	var x, y int
	var arr [NMAX]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&arr[i])
	}

	var wadah []float64

	for i := 0; i < x; i += y {
		sum := 0.0
		for j := i; j < i+y && j < x; j++ {
			sum += arr[j]
		}
		wadah = append(wadah, sum)
	}

	total := 0.0
	for i := 0; i < len(wadah); i++ {
		fmt.Print(wadah[i], " ")
		total += wadah[i]
	}
	fmt.Println()

	rata := total / float64(len(wadah))
	fmt.Println(rata)
}