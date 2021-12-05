package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

func getFileName(day int, isSample bool) string {
    prefix := "input/"

    if (isSample) {
        prefix += "sample_day_"
    } else {
        prefix += "input_day_"
    }

    return prefix + strconv.Itoa(day) + ".txt"
}

func ParseFile(day int, isSample bool) []string {
    inputFile := getFileName(day, isSample)
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

func ParseFileAndSplit(day int, isSample bool) [][]string {
    lines := ParseFile(day, isSample)
    splitLines := make([][]string, 0)

    for _, line := range lines {
        aSplitLine := strings.Split(line, " ")
        splitLines = append(splitLines, aSplitLine)
    }

    return splitLines
}

func GetInt(str string) int {
    val, err := strconv.Atoi(str)

    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }

    return val
}
