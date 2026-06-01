package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func main() {
	var x int
	data := []int{}

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		data = append(data, x)
	}

	insertionSort(data)

	for i := 0; i < len(data); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(data[i])
	}
	fmt.Println()

	tetap := true
	for i := 1; i < len(data); i++ {
		if data[i]-data[i-1] != data[1]-data[0] {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Println("Data berjarak tetap")
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}