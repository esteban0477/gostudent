package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

// Complete the formingMagicSquare function below.
func formingMagicSquare(s [][]int32) int32 {

    magicConstant := findMagicConstant(&s)

    squares := [][][]int32{
        {
            {8, 1, 6},
            {3, 5, 7},
            {4, 9, 2},
        }, {
            {6, 1, 8},
            {7, 5, 3},
            {2, 9, 4},
        }, {
            {4, 9, 2},
            {3, 5, 7},
            {8, 1, 6},
        }, {
            {2, 9, 4},
            {7, 5, 3},
            {6, 1, 8},
        }, {
            {8, 3, 4},
            {1, 5, 9},
            {6, 7, 2},
        }, {
            {4, 3, 8},
            {9, 5, 1},
            {2, 7, 6},
        }, {
            {6, 7, 2},
            {1, 5, 9},
            {8, 3, 4},
        }, {
            {2, 7, 6},
            {9, 5, 1},
            {4, 3, 8},
        },
    }

    var lowestCost int32 = 1000000000
    var cost int32

    for _, value := range squares {
        cost = 0
        for i := 0; i < len(s); i++ {
            for j := 0; j < len(s); j++ {

                if s[i][j] != value[i][j] {
                    cost += int32(math.Abs(float64(s[i][j] - value[i][j])))
                }

            }
        }

        if lowestCost > cost && checkMagicSquare(value, magicConstant) {
            lowestCost = cost
        }

    }

    return lowestCost

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    var s [][]int32
    for i := 0; i < 3; i++ {
        sRowTemp := strings.Split(readLine(reader), " ")

        var sRow []int32
        for _, sRowItem := range sRowTemp {
            sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
            checkError(err)
            sItem := int32(sItemTemp)
            sRow = append(sRow, sItem)
        }

        if len(sRow) != 3 {
            panic("Bad input")
        }

        s = append(s, sRow)
    }

    result := formingMagicSquare(s)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
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

func findMagicConstant(s *[][]int32) int32 {
    var n int32 = int32(len(*s))
    return (n * ((n*n + 1) / 2))
}
