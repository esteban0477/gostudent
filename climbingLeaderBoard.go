package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the climbingLeaderboard function below.
func climbingLeaderboard(scores []int32, alice []int32) []int32 {

    AlicePos := make([]int32, len(alice))

    for index, valAlice := range alice {
        pos := 1
        for i := 0; i < len(scores); i++ {
            if i > 0 {
                if scores[i] < scores[i-1] {
                    pos++
                }
            }
            if valAlice >= scores[i] {
                AlicePos[index] = int32(pos)
                break
            }
            if i == (len(scores) - 1) {
                AlicePos[index] = int32(pos + 1)
            }

        }
    }

    return AlicePos
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    scoresTemp := strings.Split(readLine(reader), " ")

    var scores []int32

    for i := 0; i < int(scoresCount); i++ {
        scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
        checkError(err)
        scoresItem := int32(scoresItemTemp)
        scores = append(scores, scoresItem)
    }

    aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    aliceTemp := strings.Split(readLine(reader), " ")

    var alice []int32

    for i := 0; i < int(aliceCount); i++ {
        aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
        checkError(err)
        aliceItem := int32(aliceItemTemp)
        alice = append(alice, aliceItem)
    }

    result := climbingLeaderboard(scores, alice)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
