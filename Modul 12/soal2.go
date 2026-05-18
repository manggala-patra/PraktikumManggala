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

	ketua := 1
	wakil := 2

	for i := 1; i <= 20; i++ {
		if suara[i] > suara[ketua] || (suara[i] == suara[ketua] && i < ketua) {
			wakil = ketua
			ketua = i
		} else if i != ketua && (suara[i] > suara[wakil] || (suara[i] == suara[wakil] && i < wakil)) {
			wakil = i
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}