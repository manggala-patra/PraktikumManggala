package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	rand.Seed(time.Now().UnixNano())

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		nStr := strings.TrimSpace(parts[1])
		n, err := strconv.Atoi(nStr)
		if err != nil {
			continue
		}

		inside := 0

		for i := 0; i < n; i++ {
			x := rand.Float64()*2 - 1
			y := rand.Float64()*2 - 1

			if x*x+y*y <= 1 {
				inside++
			}
		}

		pi := 4.0 * float64(inside) / float64(n)

		fmt.Println("Banyak Topping:", n)
		fmt.Println("Topping pada Pizza:", inside)
		fmt.Printf("PI : %.10f\n", pi)
		fmt.Println()
	}
}