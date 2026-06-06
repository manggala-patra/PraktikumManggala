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

func ambilKartu(d *Dominoes) Domino {
	if d.sisa <= 0 {
		return Domino{-1, -1}
	}
	k := d.kartu[d.sisa-1]
	d.sisa--
	return k
}

func gambarKartu(k Domino, sisi int) int {
	if sisi == 1 {
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

	fmt.Println("=== GAME GAPLEH SEDERHANA ===\n")

	totalPlayer := 4
	skor := make([]int, totalPlayer)

	for p := 0; p < totalPlayer; p++ {
		fmt.Printf("Player %d:\n", p+1)

		for i := 0; i < 4; i++ {
			k := ambilKartu(&deck)
			fmt.Printf("  Kartu: [%d|%d] nilai=%d\n",
				gambarKartu(k, 1), gambarKartu(k, 2), nilaiKartu(k))
			skor[p] += nilaiKartu(k)
		}

		fmt.Printf("  Total nilai Player %d = %d\n\n", p+1, skor[p])
	}

	menang := 0
	for i := 1; i < totalPlayer; i++ {
		if skor[i] > skor[menang] {
			menang = i
		}
	}

	fmt.Println("Pemenang adalah Player", menang+1, "dengan skor", skor[menang])
}