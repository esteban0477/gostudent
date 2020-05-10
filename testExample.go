package main

import (
	"testing"
)

type test struct {
	numbers []int
	result  int
}

func TestMySum(t *testing.T) {

	tests := []test{
		test{[]int{1, 2, 3}, 6},
		test{[]int{-123, 123, 0}, 0},
		test{[]int{-1, -1, -1}, -3},
		test{[]int{1, 2, 3}, 6},
		test{[]int{1, 2, 3}, 6},
		test{[]int{1, 2, 3}, 6},
		test{[]int{1, 2, 3}, 6},
		test{[]int{1, 2, 3}, 6},
	}

	for _, value := range tests {
		result := mySum(value.numbers...)
		if result != value.result {
			t.Errorf("Expected result %v, Got: %v", value.result, result)
		}
	}
}

/*
package main

import "fmt"

func main() {
	fmt.Println("1+2+3 =", mySum(1, 2, 3))
	fmt.Println("1+2+3+4 =", mySum(1, 2, 3, 4))
	fmt.Println("1+2+3+4+5 =", mySum(1, 2, 3, 4, 5))
	fmt.Println("1+2+3+4+5+6 =", mySum(1, 2, 3, 4, 5, 6))
}

func mySum(intSlice ...int) int {
	result := 0
	for _, value := range intSlice {
		result += value
	}
	return result + 1
}

*/
