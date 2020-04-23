package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x [500]int
	for i := 0; i <= len(x)-1; i++ {
		x[i] = rand.Intn(100)
		fmt.Printf("%v\t", x[i])
	}
	fmt.Println()

	/* Now, as a task, let's sort the array from the lower to the higher */
	for i := 0; i < (len(x)); i++ {
		for j := i; j < (len(x)); j++ {
			if x[j] < x[i] {
				x[j] = x[j] + x[i]
				x[i] = x[j] - x[i]
				x[j] = x[j] - x[i]
			}
		}
	}
	fmt.Println("Let's print the array from the lowest to the highest value")
	for i := 0; i < len(x); i++ {
		fmt.Printf("%v\t", x[i])
	}
	fmt.Println()

	/* Now, as a task, let's sort the array from the higher to the lower */
	for i := 0; i < (len(x)); i++ {
		for j := i; j < (len(x)); j++ {
			if x[j] > x[i] {
				x[j] = x[j] + x[i]
				x[i] = x[j] - x[i]
				x[j] = x[j] - x[i]
			}
		}
	}
	fmt.Println("Let's print the array from the highest to the lowest value")
	for i := 0; i < len(x); i++ {
		fmt.Printf("%v\t", x[i])
	}
	fmt.Println()

	/* Another task: count the number of repetions per each number a write it to the std output */
	var values [len(x)]int
	var repetitions [len(x)]int
	for i := 0; i < (len(values)); i++ {
		values[i] = x[i]
		for j := 0; j < (len(values)); j++ {
			if values[i] == x[j] {
				repetitions[i]++
			}
		}
	}
	fmt.Println("Let's print the number of ocurrences")
	for i := 0; i < (len(values) - 1); i++ {
		if values[i] != values[i+1] {
			fmt.Printf("Number: %v was seen %v times on the array\n", values[i], repetitions[i])
		}
	}
}
