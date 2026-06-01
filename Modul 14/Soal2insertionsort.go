package main

import "fmt"

const NMAX = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku struct {
	data []Buku
}

func InputBuku(P *DaftarBuku, n int) {
	P.data = make([]Buku, n)

	for i := 0; i < n; i++ {
		fmt.Scan(
			&P.data[i].id,
			&P.data[i].judul,
			&P.data[i].penulis,
			&P.data[i].penerbit,
			&P.data[i].eksemplar,
			&P.data[i].tahun,
			&P.data[i].rating,
		)
	}
}

func CetakTerfavorit(P DaftarBuku, n int) {
	if n > len(P.data) {
		n = len(P.data)
	}

	for i := 0; i < n; i++ {
		fmt.Println(
			P.data[i].id,
			P.data[i].judul,
			P.data[i].penulis,
			P.data[i].penerbit,
			P.data[i].eksemplar,
			P.data[i].tahun,
			P.data[i].rating,
		)
	}
}

func UrutBuku(P *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		temp := P.data[i]
		j := i - 1

		for j >= 0 && P.data[j].rating < temp.rating {
			P.data[j+1] = P.data[j]
			j--
		}

		P.data[j+1] = temp
	}
}

func CetakTerbaru(P DaftarBuku, n int) {
	if n > len(P.data) {
		n = len(P.data)
	}

	for i := 0; i < n; i++ {
		fmt.Println(
			P.data[i].id,
			P.data[i].judul,
			P.data[i].penulis,
			P.data[i].penerbit,
			P.data[i].eksemplar,
			P.data[i].tahun,
			P.data[i].rating,
		)
	}
}

func CariBuku(P DaftarBuku, n int, rating int) {
	left := 0
	right := n - 1
	idx := -1

	for left <= right {
		mid := (left + right) / 2

		if P.data[mid].rating == rating {
			idx = mid
			break
		} else if P.data[mid].rating < rating {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if idx == -1 {
		fmt.Println("Buku tidak ditemukan")
	} else {
		fmt.Println(
			P.data[idx].id,
			P.data[idx].judul,
			P.data[idx].penulis,
			P.data[idx].penerbit,
			P.data[idx].eksemplar,
			P.data[idx].tahun,
			P.data[idx].rating,
		)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var pustaka DaftarBuku

	InputBuku(&pustaka, n)

	UrutBuku(&pustaka, n)

	CetakTerfavorit(pustaka, 5)

	CetakTerbaru(pustaka, 5)

	var ratingCari int
	fmt.Scan(&ratingCari)

	CariBuku(pustaka, n, ratingCari)
}