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

func SplitIntoInt(str string, char string) []int {
    var splitInt []int
    var splitStr []string

    if char == " " {
        splitStr = strings.Fields(str)
    } else {
        splitStr = strings.Split(str, char)
    }

    for _, term := range splitStr {
        splitInt = append(splitInt, GetInt(term))
    }

    return splitInt
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

func GetSring(i int) string {
    // val, err := strconv.Itoa(i)
    return strconv.Itoa(i)

    // if err != nil {
    //     // handle error
    //     fmt.Println(err)
    //     os.Exit(2)
    // }

    // return val
}

func GetIntFromBinary(str string) int {
    val, err := strconv.ParseInt(str, 2, 64)

    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }

    return int(val)
}

func GetCharAt(str string, i int) string {
    rn := []rune(str)
    return string(rn[i])
}

func PrintIntArr(arr [][]int) {
    for _, row := range arr {
        fmt.Println(row)
    }
    fmt.Println("")
}

func GetMinMax(num1 int, num2 int) (int, int) {
    min, max := -1, -1

    if (num1 > num2) {
        min = num2
        max = num1
    } else if (num1 < num2) {
        min = num1
        max = num2
    }

    return min, max
}

func GetMin(arr []int) int {
    min := arr[0]

    for _, val := range arr {
        if min > val {
            min = val
        }
    }

    return min
}

func GetMax(arr []int) int {
    max := arr[0]

    for _, val := range arr {
        if max < val {
            max = val
        }
    }

    return max
}

func Abs(num int) int {
    if (num < 0) {
        return -num
    }
    return num
}

func In(char rune, str string) bool {
    for _, char2 := range str {
        // fmt.Println(char, char2)

        if char == char2 {
            return true
        }
    }

    return false
}

func RemoveStringAtIndex(s []string, index int) []string {
    return append(s[:index], s[index+1:]...)
}

func Use(vals ...interface{}) {
    for _, val := range vals {
        _ = val
    }
}

func GetLastChar(str string) string {
    return string(str[len(str) - 1])
}
