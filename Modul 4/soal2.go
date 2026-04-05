package main

import "fmt"

func hitungSkor(skor *int, total *int) {
	*skor = 0
	*total = 0

	var t int
	for i := 0; i < 8; i++ {
		fmt.Scan(&t)
		if t <= 300 {
			*skor++
			*total += t
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	maxSkor := -1
	minWaktu := 1000000000
	pemenang := ""

	for i := 0; i < n; i++ {
		var nama string
		fmt.Scan(&nama)

		var skor, total int
		hitungSkor(&skor, &total)

		if skor > maxSkor || (skor == maxSkor && total < minWaktu) {
			maxSkor = skor
			minWaktu = total
			pemenang = nama
		}
	}

	fmt.Println(pemenang, maxSkor, minWaktu)
}