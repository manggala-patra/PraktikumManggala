package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var idx int
	fmt.Scan(&idx)
	arr = append(arr[:idx], arr[idx+1:]...)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	avg := float64(sum) / float64(len(arr))
	fmt.Println(avg)

	var total float64
	for i := 0; i < len(arr); i++ {
		total += math.Pow(float64(arr[i])-avg, 2)
	}
	std := math.Sqrt(total / float64(len(arr)))
	fmt.Println(std)

	var cari, count int
	fmt.Scan(&cari)
	for i := 0; i < len(arr); i++ {
		if arr[i] == cari {
			count++
		}
	}
	fmt.Println(count)
}