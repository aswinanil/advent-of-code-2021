package main

import (
    "fmt"
)

func getMinOrZero(arr []int) int {
    min := arr[0]

    for _, val := range arr {
        if val == 0 {
            return 0
        }

        if min > val {
            min = val
        }
    }

    return min
}

func passOneDay(fishes []int) ([]int, int) {
    newFishes := make([]int, 0)

    min := getMinOrZero(fishes)

    if (min != 0) {
        for i, _ := range fishes {
            fishes[i] -= min
        }

        return fishes, min
    }


    for i, _ := range fishes {
        if fishes[i] == 0 {
            fishes[i] = 6
            newFishes = append(newFishes, 8)
        } else {
            fishes[i]--
        }
    }

    return append(fishes, newFishes...), 1
}

func passXDays(fishes []int, days int) []int {
    i := 0
    var daysPassed int

    for i < days {
        fishes, daysPassed = passOneDay(fishes)
        i += daysPassed
    }

    return fishes
}

func main() {
    lines := ParseFile(6, true)
    fishes := SplitIntoInt(lines[0], ",")

    fishes = passXDays(fishes, 80)  // Part 1
    fmt.Println(len(fishes))
}
