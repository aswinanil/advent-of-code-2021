package main

import (
    "fmt"
)

func checkIncrements(lines []int) int {
    prevValue := -1
    incrementCount := 0

    for _, val := range lines {
        if prevValue != -1 && val > prevValue {
            incrementCount += 1
        }

        prevValue = val
    }

    return incrementCount
}

func getIntListFromStringList(lines []string) []int {
    intList := make([]int, 0)

    for _, val := range lines {
        intVal := GetInt(val)
        intList = append(intList, intVal)
    }

    return intList
}

func getThreeNumWindow(lines []int) []int {
    windowSums := make([]int, 0)

    for i := 0; i < len(lines) - 2; i++ {
        windowSums = append(windowSums, lines[i] + lines[i+1] + lines[i+2])
    }

    return windowSums
}

func main() {
    lines := ParseFile(1, false)
    intList := getIntListFromStringList(lines)

    // Part 1
    fmt.Println(checkIncrements(intList))

    // Part 2
    windowSums := (getThreeNumWindow(intList))
    fmt.Println(checkIncrements(windowSums))
}
