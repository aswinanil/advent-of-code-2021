package main

import (
    "fmt"
    "sync"
    "strings"
)

/*
For part 2, these digits denote position of the signal's char

 1111
2    3
2    3
 4444
5    6
5    6
 7777

*/

func subtractString(anchor string, toSubtract string) string {
    var diff string

    for _, char := range anchor {
        if !In(char, toSubtract) {
            diff += string(char)
        }
    }

    return diff
}

func getSignalNumberPairings(input []string) [10]string {
    var one, four, seven, eight, six, nine, zero, two, three, five string  // In order of their values being discovered
    var fiveCharSignals, sixCharSignals []string

    for _, signal := range input {
        l := len(signal)
        switch l {
        case 2:
            one = signal
        case 3:
            seven = signal
        case 4:
            four = signal
        case 7:
            eight = signal
        case 5:
            fiveCharSignals = append(fiveCharSignals, signal)
        case 6:
            sixCharSignals = append(sixCharSignals, signal)
        }
    }

    // pos1 := subtractString(four, one)

    for i, signal := range sixCharSignals {
        diff := subtractString(signal, one)

        if len(signal) - len(diff) == 1 {
            six = signal  // Only six, amongst 6 digit signals, is missing both digits found in one
            sixCharSignals = RemoveStringAtIndex(sixCharSignals, i)
            break
        }
    }

    pos3 := subtractString(one, six)
    pos6 := subtractString(one, pos3)

    for i, signal := range sixCharSignals {
        diff := subtractString(signal, four)

        if len(diff) == 2 {
            nine = signal  // Only nine, amongst 6 digit signals, overlap completely with four
            RemoveStringAtIndex(sixCharSignals, i)
            break
        }
    }

    zero = sixCharSignals[0]
    // pos4 := subtractString(eight, zero)

    for i, signal := range fiveCharSignals {
        if (!In(rune(pos6[0]), signal)) {
            two = signal  // Only two, amongst 5 digit signals, doesn't contain pos6 char
            fiveCharSignals = RemoveStringAtIndex(fiveCharSignals, i)
            break
        }
    }

    if In(rune(pos3[0]), fiveCharSignals[0]) {
        three = fiveCharSignals[0]
        five = fiveCharSignals[1]
    } else {
        three = fiveCharSignals[1]
        five = fiveCharSignals[0]
    }

    return [10]string{zero, one, two, three, four, five, six, seven, eight, nine}
}

func decodeSignal(decodedSignals[10]string, signal string) string {
    for i, decodedSignal := range decodedSignals {
        diff := subtractString(signal, decodedSignal)

        if len(decodedSignal) == len(signal) && len(diff) == 0 {
            return GetSring(i)
        }
    }

    return "0"
}

func getLineOutputSum(input []string, output []string, sumChan chan int) {
    sum := ""

    decodedSignals := getSignalNumberPairings(input)

    for _, signal := range output {
        sum += decodeSignal(decodedSignals, signal)
    }

    sumChan <- GetInt(sum)
}

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

func getTotal(inputs [][]string, outputs [][]string, wg *sync.WaitGroup, part int) {
    totalChan := make(chan int, len(outputs))

    for i, output := range outputs {
        if part == 1 {
            go getLineUniqueCount(output, totalChan)
        } else {
            go getLineOutputSum(inputs[i], output, totalChan)
        }
    }

    total := 0
    for i:=0; i<len(outputs); i++ {
        total += <- totalChan
    }

    fmt.Printf("Part %d: %d\n", part, total)
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
    inputs, outputs := parseSignals(lines)

    var wg sync.WaitGroup
    wg.Add(2)

    go getTotal(inputs, outputs, &wg, 1)
    go getTotal(inputs, outputs, &wg, 2)

    wg.Wait()
}
