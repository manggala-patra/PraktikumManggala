package main

import (
	"fmt"
)

const NMAX = 127

type tabel struct {
	data [NMAX]rune
	n    int
}

func isiArray(t *tabel, n int) {
	t.n = n
	for i := 0; i < n; i++ {
		fmt.Scanf(" %c", &t.data[i])
	}
}

func cetakArray(t tabel) {
	for i := 0; i < t.n; i++ {
		fmt.Printf("%c ", t.data[i])
	}
	fmt.Println()
}

func balikArray(t *tabel) {
	for i := 0; i < t.n/2; i++ {
		t.data[i], t.data[t.n-1-i] = t.data[t.n-1-i], t.data[i]
	}
}

func palindrom(t tabel) bool {
	for i := 0; i < t.n/2; i++ {
		if t.data[i] != t.data[t.n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	fmt.Scan(&n)
	isiArray(&tab, n)

	cetakArray(tab)

	balikArray(&tab)
	cetakArray(tab)

	if palindrom(tab) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}