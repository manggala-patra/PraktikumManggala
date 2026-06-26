package main

import "fmt"

func main() {
	var suara [21]int
	var x int
	total := 0
	valid := 0

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		total++

		if x >= 1 && x <= 20 {
			suara[x]++
			valid++
		}
	}

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", valid)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}