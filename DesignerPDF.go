package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the designerPdfViewer function below.
func designerPdfViewer(h []int32, word string) int32 {

    intSlice := []byte(word)
    wordHeight := make([]int32, len(intSlice))

    for index, value := range intSlice {
        intSlice[index] = value - 97
        wordHeight[index] = h[int(intSlice[index])]
    }

    highest := wordHeight[0]
    for _, value := range wordHeight {
        if highest < value {
            highest = value
        }
    }

    return (highest * int32(len(intSlice)))
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    hTemp := strings.Split(readLine(reader), " ")

    var h []int32

    for i := 0; i < 26; i++ {
        hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
        checkError(err)
        hItem := int32(hItemTemp)
        h = append(h, hItem)
    }

    word := readLine(reader)

    result := designerPdfViewer(h, word)

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
