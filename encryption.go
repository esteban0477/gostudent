package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "math"
)

// Complete the encryption function below.
func encryption(s string) string {

    byteSlice := []byte(s)
    rows := int(math.Floor(math.Sqrt(float64(len(byteSlice)))))
    columns := int(math.Ceil(math.Sqrt(float64(len(byteSlice)))))

    for !((rows * columns) >= len(byteSlice)) {
        rows++
    }

    matrix := make([][]byte, rows)
    for i := 0; i < rows; i++ {
        matrix[i] = make([]byte, columns)
    }

    counterLen := 0
    for i := 0; i < rows; i++ {
        for j := 0; j < columns; j++ {
            if counterLen < len(byteSlice) {
                matrix[i][j] = byteSlice[counterLen]
                counterLen++
            }
        }
    }

    outSlice := make([]byte, len(byteSlice)+columns)
    counter := 0

    for i := 0; i < columns; i++ {
        for j := 0; j < rows; j++ {
            if counter < len(outSlice) && matrix[j][i] != 0 {
                outSlice[counter] = matrix[j][i]
                counter++
            }
        }
        if counter < len(outSlice) {
            outSlice[counter] = byte(32)
            counter++
        }
    }

    return string(outSlice)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    result := encryption(s)

    fmt.Fprintf(writer, "%s\n", result)

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
