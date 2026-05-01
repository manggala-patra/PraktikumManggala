package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Scan(&klubA)
	fmt.Scan(&klubB)

	pemenang := []string{}
	no := 1

	for {
		var skorA, skorB int
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}

		no++
	}

	for i := 0; i < len(pemenang); i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, pemenang[i])
	}

	fmt.Println("Pertandingan selesai")
}