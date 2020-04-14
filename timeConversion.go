package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
    output := make([]byte, len(s) - 2 )

    if s[8] == 80 {

        hour := ((s[0] - 48) * 10) + (s[1] - 48) + 12

        if hour == 24 {
            hour = 12
        }
        
        output[0] = (hour / 10) + 48
        output[1] = (hour % 10) + 48

        for i := 2; i < len(s)-2; i++ {
            output[i] = s[i]
        }

    } else if s[8] == 65 {

        hour := ((s[0] - 48) * 10) + (s[1] - 48) + 12
        
        if hour == 24 {
            hour = 0
            output[0] = (hour / 10) + 48
            output[1] = (hour % 10) + 48
            for i := 2; i < len(s)-2; i++ {
                output[i] = s[i]
            }
        } else {
            for i := 0; i < len(s)-2; i++ {
                output[i] = s[i]
            }
        }
    }
    return(string(output))
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer outputFile.Close()

    writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

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
