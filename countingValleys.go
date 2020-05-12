package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {

    var lastlevel, level, valley, mountain int32 = 0, 0, 0, 0
    xbyte := []byte(s)
    
    for _, value := range xbyte {
        switch string(value) {
        case "U":
            level++
        case "D":
            level--
        }

        if level < 0 && lastlevel == 0 {
            valley++
        }

        if level > 0 && lastlevel == 0 {
            mountain++
        }

        lastlevel = level

    }

    return valley

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    s := readLine(reader)

    result := countingValleys(n, s)

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
