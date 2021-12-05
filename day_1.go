package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
)

// var inputFile = "input/sample_day_1.txt"
var inputFile = "input/input_day_1.txt"

func parseFile() []string {
    lines := make([]string, 0)

    file, err := os.Open(inputFile)
    if err != nil {
        fmt.Println(err)
        return lines
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    return lines
}

func checkIncrements(lines []string) int {

    prevValue := -1
    incrementCount := 0

    for _, val := range lines {
        val, err := strconv.Atoi(val)

        if err != nil {
            // handle error
            fmt.Println(err)
            os.Exit(2)
        }

        if prevValue != -1 && val > prevValue {
            incrementCount += 1
        }

        prevValue = val
    }

    return incrementCount
}

func main() {
    lines := parseFile()
    fmt.Println(checkIncrements(lines))
}
