package main

import "fmt"

func main() {
	var angka float64
	var jumlah float64
	var banyakData int

	fmt.Println("Masukkan bilangan (akhiri dengan 9999):")

	for {
		fmt.Scan(&angka)

		if angka == 9999 {
			break
		}

		jumlah += angka
		banyakData++
	}

	if banyakData > 0 {
		rataRata := jumlah / float64(banyakData)
		fmt.Printf("Rata-rata = %.2f\n", rataRata)
	} else {
		fmt.Println("Tidak ada data yang dimasukkan.")
	}
}