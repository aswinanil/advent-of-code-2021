package main

import (
    "fmt"
    "sync"
    "strings"
)

func isUnique(signal string) bool {
    l := len(signal)

    switch l {
    case 2, 3, 4, 7:
        return true
    default:
        return false
    }
}

func getLineUniqueCount(output []string, countChan chan int) {
    count := 0

    for _, signal := range output {
        if isUnique(signal) {
            count++
        }
    }

    countChan <- count
}

func getUniqueCount(outputs [][]string, wg *sync.WaitGroup) {
    countChan := make(chan int, len(outputs))

    for _, output := range outputs {
        go getLineUniqueCount(output, countChan)
    }

    count := 0
    for i:=0; i<len(outputs); i++ {
        count += <- countChan
    }

    fmt.Printf("Part %d: %d\n", 1, count)
    defer wg.Done()
}

func parseSignals(lines []string) ([][]string, [][]string) {
    inputs := make([][]string, 0)
    outputs := make([][]string, 0)

    for _, line := range lines {
        signals := strings.Split(line, "|")
        inputs = append(inputs, strings.Split(signals[0], " "))
        outputs = append(outputs, strings.Split(signals[1], " "))
    }

    return inputs, outputs
}

func main() {
    lines := ParseFile(8, false)
    _, outputs := parseSignals(lines)

    var wg sync.WaitGroup
    wg.Add(1)

    go getUniqueCount(outputs, &wg)

    wg.Wait()
}
