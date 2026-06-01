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

		if x == -5313 {
			break
		}

		if x == 0 {
			if len(data) > 0 {
				insertionSort(data)

				n := len(data)
				var median int

				if n%2 == 1 {
					median = data[n/2]
				} else {
					median = (data[n/2-1] + data[n/2]) / 2
				}

				fmt.Println(median)
			}
		} else {
			data = append(data, x)
		}
	}
}