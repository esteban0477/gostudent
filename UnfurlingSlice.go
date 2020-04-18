package main

import "fmt"

func main() {
	intSlice := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	totalSum := foo(intSlice...)
	fmt.Println(totalSum)
	totalSum = bar(intSlice)
	fmt.Println(totalSum)
}

func foo(Slicepar ...int) int {
	fmt.Printf("The type of Slicepar is %T\n", Slicepar)
	total := 0
	for _, value := range Slicepar {
		total += value
	}
	return total
}

func bar(Slicepar []int) int {
	fmt.Printf("The type of Slicepar is %T\n", Slicepar)
	total := 0
	for _, value := range Slicepar {
		total += value
	}
	return total
}
