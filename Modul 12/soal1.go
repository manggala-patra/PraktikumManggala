package main

import "fmt"

func main() {
	var angka int
	totalMasuk := 0
	totalSah := 0
	suara := make([]int, 21)

	for {
		fmt.Scan(&angka)

		if angka == 0 {
			break
		}

		totalMasuk++

		if angka >= 1 && angka <= 20 {
			suara[angka]++
			totalSah++
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalSah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}