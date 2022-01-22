package main

import (
    "fmt"
)

func passOneDay(histogram map[int]int) map[int]int {
    newHistogram := make(map[int]int)

    for i:=8; i>0; i-- {
        newHistogram[i-1] = histogram[i]
    }

    newHistogram[8] = histogram[0]
    newHistogram[6] += histogram[0]

    return newHistogram
}

func passXDays(histogram map[int]int, days int) map[int]int {
    for i:=0; i<days; i++ {
        histogram = passOneDay(histogram)
    }
    return histogram
}

func getHistogram(timers []int) map[int]int {
    timerCountMap := make(map[int]int)

    for _, timer := range timers {
        timerCountMap[timer]++
    }
    return timerCountMap
}

func getTotalCount(histogram map[int]int) int {
    totalCount := 0

    for _, count := range histogram {
        totalCount += count
    }
    return totalCount
}

func main() {
    lines := ParseFile(6, false)
    timers := SplitIntoInt(lines[0], ",")
    histogram := getHistogram(timers)

    histogram = passXDays(histogram, 80)  // Part 1
    fmt.Println(getTotalCount(histogram))

    histogram = passXDays(histogram, 256-80)  // Part 2
    fmt.Println(getTotalCount(histogram))
}
