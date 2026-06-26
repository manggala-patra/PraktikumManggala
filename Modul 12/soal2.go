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

	ketua := 1
	wakil := 2

	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] || (suara[i] == suara[ketua] && i < ketua) {
			ketua = i
		}
	}

	for i := 1; i <= 20; i++ {
		if i == ketua {
			continue
		}

		if suara[i] > suara[wakil] ||
			(suara[i] == suara[wakil] && i < wakil) ||
			wakil == ketua {

			if suara[i] < suara[ketua] || (suara[i] == suara[ketua] && i > ketua) {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", valid)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}