package main

import (
	"fmt"
	"math/rand"
)

func main() {

	s := [][]int32{
		{6, 1, 8},
		{7, 5, 3},
		{2, 9, 4},
	}

	magicConstant := findMagicConstant(&s)

	/*
		var sumRow int32
		var sumColumn int32
		var sumBackSlash int32
		var sumSlash int32
	*/

	/*squares := [][][]int32{
		{
			{6, 1, 8},
			{7, 5, 3},
			{2, 9, 4},
		}, {
			{6, 1, 8},
			{7, 5, 3},
			{2, 9, 4},
		},
	}*/

	candidate := make([][]int32, len(s))
	for i := 0; i < len(candidate); i++ {
		candidate[i] = make([]int32, len(candidate))
	}

	// magicSquares := 8
	// counter := 0
	var breaker bool = true
	var temporal int32

	for breaker {

		for x := 0; x < len(candidate); x++ {
			for y := 0; y < len(candidate); y++ {

				if x == y && x != 1 {
					temporal = rand.Int31n(9) + 1
					if temporal%2 == 0 {
						for a := 0; a < len(candidate); a++ {
							for b := 0; b < len(candidate); b++ {
								if candidate[a][b] == temporal {
									temporal = 0
								}
							}
						}
						if temporal != 0 {
							candidate[x][y] = temporal
						}
					}
				}

				if x == y && x == 1 {
					candidate[x][y] = int32(5)
				}

				if (x+y) == (len(candidate)-1) && x != 1 {
					temporal = rand.Int31n(9) + 1
					if temporal%2 == 0 {
						for a := 0; a < len(candidate); a++ {
							for b := 0; b < len(candidate); b++ {
								if candidate[a][b] == temporal {
									temporal = 0
								}
							}
						}
						if temporal != 0 {
							candidate[x][y] = temporal
						}
					}
				}

				if (x+y) != (len(candidate)-1) && x != y {
					temporal = rand.Int31n(9) + 1
					if temporal%2 != 0 {
						for a := 0; a < len(candidate); a++ {
							for b := 0; b < len(candidate); b++ {
								if candidate[a][b] == temporal {
									temporal = 0
								}
							}
						}
						if temporal != 0 {
							candidate[x][y] = temporal
						}
					}
				}

			}
		}

		breaker = false

		for x := 0; x < len(candidate); x++ {
			for y := 0; y < len(candidate); y++ {
				if candidate[x][y] == 0 {
					breaker = true
				}
			}
		}

		if !breaker {
			breaker = !(checkMagicSquare(candidate, magicConstant))
		}
	}

	fmt.Println(candidate)

	// fmt.Println(checkMagicSquare(s, magicConstant))

}

func findMagicConstant(s *[][]int32) int32 {
	var n int32 = int32(len(*s))
	return (n * ((n*n + 1) / 2))
}

func checkMagicSquare(s [][]int32, magicConstant int32) bool {
	var sumRow int32
	var sumColumn int32
	var sumBackSlash int32
	var sumSlash int32

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			sumRow += s[i][j]
			sumColumn += s[j][i]

			if i == j {
				sumBackSlash += s[i][j]
			}

			if (i + j) == (len(s) - 1) {
				sumSlash += s[i][j]
			}
		}

		if sumRow != magicConstant || sumColumn != magicConstant {
			return false
		}

		sumRow = 0
		sumColumn = 0
	}

	if sumBackSlash != magicConstant || sumSlash != magicConstant {
		return false
	}

	return true
}
