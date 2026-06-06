package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	suit1 int
	suit2 int
	balak bool
}

type Dominoes struct {
	kartu []Domino
	sisa  int
}

func buatDomino() Dominoes {
	var d []Domino

	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			b := false
			if i == j {
				b = true
			}
			d = append(d, Domino{i, j, b})
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

func ambilKartu(d *Dominoes) Domino {
	if d.sisa <= 0 {
		return Domino{-1, -1, false}
	}

	k := d.kartu[d.sisa-1]
	d.sisa--
	return k
}

func gambarKartu(k Domino, suit int) int {
	if suit == 1 {
		return k.suit1
	}
	return k.suit2
}

func nilaiKartu(k Domino) int {
	return k.suit1 + k.suit2
}

func main() {
	deck := buatDomino()

	kocokKartu(&deck)

	fmt.Println("Ambil 5 kartu domino:\n")

	for i := 0; i < 5; i++ {
		k := ambilKartu(&deck)
		fmt.Printf("Kartu %d: [%d|%d] | nilai: %d\n",
			i+1, gambarKartu(k, 1), gambarKartu(k, 2), nilaiKartu(k))
	}

	fmt.Println("\nSisa kartu:", deck.sisa)
}