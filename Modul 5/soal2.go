package main

import "fmt"

func cetakBintang(n int, i int) {
	if i > n {
		return
	}

	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	cetakBintang(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	cetakBintang(n, 1)
}