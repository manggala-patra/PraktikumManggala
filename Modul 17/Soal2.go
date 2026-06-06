package main

import "fmt"

func main() {
	var x string
	var n int

	fmt.Scan(&x)
	fmt.Scan(&n)

	jumlahKemunculan := 0
	posisiPertama := -1

	for i := 0; i < n; i++ {
		var data string
		fmt.Scan(&data)

		if data == x {
			jumlahKemunculan++

			if posisiPertama == -1 {
				posisiPertama = i + 1
			}
		}
	}

	if jumlahKemunculan > 0 {
		fmt.Println("String ditemukan")
	} else {
		fmt.Println("String tidak ditemukan")
	}

	if posisiPertama != -1 {
		fmt.Println(posisiPertama)
	} else {
		fmt.Println("Tidak ditemukan")
	}

	fmt.Println(jumlahKemunculan)

	if jumlahKemunculan >= 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}