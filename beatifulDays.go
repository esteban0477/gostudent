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

// Complete the beautifulDays function below.
func beautifulDays(i int32, j int32, k int32) int32 {
    counter := int32(0)
    for ; i <= j; i++ {
        if (int32(math.Abs(float64(i-backwardsNumber(i)))) % k) == 0 {
            counter++
        }
    }
    return counter
}

func backwardsNumber(value int32) int32 {
    digits := 1
    for i := value / 10; i >= 1; digits++ {
        i = i / 10
    }
    var iback int32
    for divPow := 1; divPow <= digits; divPow++ {
        iback += (value % 10) * int32(math.Pow10(digits-divPow))
        value /= 10
    }
    return iback
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    ijk := strings.Split(readLine(reader), " ")

    iTemp, err := strconv.ParseInt(ijk[0], 10, 64)
    checkError(err)
    i := int32(iTemp)

    jTemp, err := strconv.ParseInt(ijk[1], 10, 64)
    checkError(err)
    j := int32(jTemp)

    kTemp, err := strconv.ParseInt(ijk[2], 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := beautifulDays(i, j, k)

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
