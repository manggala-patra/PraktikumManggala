package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	var jumlah float64
	var piLama, piBaru float64
	var ketemu bool
	var indeks int

	for i := 1; i <= n; i++ {
		suku := 1.0 / float64(2*i-1)

		if i%2 == 0 {
			suku = -suku
		}

		jumlah += suku
		piBaru = 4 * jumlah

		if i > 1 {
			if math.Abs(piBaru-piLama) < 0.00001 && !ketemu {
				ketemu = true
				indeks = i

				fmt.Printf("Hasil PI: %.9f\n", piLama)
				fmt.Printf("Hasil PI: %.9f\n", piBaru)
				fmt.Printf("Pada i ke: %d\n", indeks)
			}
		}

		piLama = piBaru
	}

	if !ketemu {
		fmt.Printf("Hasil PI: %.9f\n", piBaru)
	}
}