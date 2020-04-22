package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {
    months := map[int32]int32{
        1:  31,
        2:  28,
        3:  31,
        4:  30,
        5:  31,
        6:  30,
        7:  31,
        8:  31,
        9:  30,
        10: 31,
        11: 30,
        12: 31,
    }

    switch {
    case year >= 1700 && year <= 1917:
        if year%4 == 0 {
            months[2] = 29
        }

    case year == 1918:
        months[2] = 15

    case year >= 1919 && year <= 2700:
        if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
            months[2] = 29
        }
    }

    var sumDays, day int32
    var month int32

    monthsChar := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
    daysChar := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12",
        "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}

    for i := int32(1); i <= int32(len(months)); i++ {
        sumDays += months[i]
        if sumDays >= 256 {
            sumDays -= months[i]
            day = 256 - sumDays
            month = i
            break
        }
    }

    result := daysChar[day-1] + "." + monthsChar[month-1] + "." + strconv.FormatInt(int64(year), 10)
    return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    year := int32(yearTemp)

    result := dayOfProgrammer(year)

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
