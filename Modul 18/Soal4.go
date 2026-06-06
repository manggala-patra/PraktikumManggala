package main

import (
	"fmt"
)

var text string
var idx int

func start(input string) {
	text = input
	idx = 0
}

func maju() {
	if idx < len(text) {
		idx++
	}
}

func eop() bool {
	return idx >= len(text) || text[idx] == '.'
}

func cc() byte {
	return text[idx]
}

func main() {
	input := "LEARN TO PROGRAM WITH LEARNING LE. "

	start(input)

	totalChar := 0
	countA := 0
	countLE := 0

	for !eop() {
		c := cc()

		totalChar++

		if c == 'A' {
			countA++
		}

		if idx < len(text)-1 {
			if text[idx] == 'L' && text[idx+1] == 'E' {
				countLE++
			}
		}

		maju()
	}

	var freqA float64
	if totalChar > 0 {
		freqA = float64(countA) / float64(totalChar)
	}

	fmt.Println("Total karakter:", totalChar)
	fmt.Println("Jumlah huruf A:", countA)
	fmt.Println("Frekuensi A:", freqA)
	fmt.Println("Jumlah kata LE:", countLE)
}