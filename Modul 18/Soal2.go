package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	suit1 int
	suit2 int
}

type Dominoes struct {
	kartu []Domino
	sisa  int
}

func buatDomino() Dominoes {
	var d []Domino

	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			d = append(d, Domino{i, j})
		}
	}

	return Dominoes{d, len(d)}
}

func kocokKartu(d *Dominoes) {
	rand.Seed(time.Now().UnixNano())
	for i := range d.kartu {
		j := rand.Intn(len(d.kartu))
		d.kartu[i], d.kartu[j] = d.kartu[j], d.kartu[i]
	}
}

func galiKartu(d *Dominoes, target Domino) Domino {
	for d.sisa > 0 {
		k := d.kartu[d.sisa-1]
		d.sisa--

		if k.suit1 == target.suit1 || k.suit1 == target.suit2 ||
			k.suit2 == target.suit1 || k.suit2 == target.suit2 {
			return k
		}
	}
	return Domino{-1, -1}
}

func sepasangKartu(a Domino, b Domino) bool {
	total := a.suit1 + a.suit2 + b.suit1 + b.suit2
	return total == 12
}

func main() {
	deck := buatDomino()
	kocokKartu(&deck)

	target := Domino{3, 5}

	kartu := galiKartu(&deck, target)
	fmt.Println("Kartu ditemukan:", kartu)

	a := Domino{6, 3}
	b := Domino{2, 1}

	fmt.Println("Sepasang valid (total 12):", sepasangKartu(a, b))
}