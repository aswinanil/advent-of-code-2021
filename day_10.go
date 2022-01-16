package main

import (
    "fmt"
    "sort"
)

type results struct {
    char string
    isCorrupt bool
    balance string
}

func isOpener(char rune) bool {
    openers := "([{<"
    return In(char, openers)
}

func isCloser(char rune) bool {
    closers := ")]}>"
    return In(char, closers)
}

func getMatch(char rune) string {
    switch string(char) {
    case "(":
        return ")"
    case ")":
        return "("
    case "[":
        return "]"
    case "]":
        return "["
    case "{":
        return "}"
    case "}":
        return "{"
    case "<":
        return ">"
    case ">":
        return "<"
    }
    return ""
}

func findFirstIncorrectChar(line string, resultsChan chan results) {
    var openers, incorretChar string

    for _, char := range line {
        if isOpener(char) {
            openers = openers + string(char)
        } else if isCloser(char) {
            if GetLastChar(openers) != getMatch(char) {
                incorretChar = string(char)
                break
            } else {
                openers = openers[:len(openers)-1]
            }
        }
    }

    if (len(incorretChar) == 0) {
        resultsChan <- results{"", false, openers}
    } else {
        resultsChan <- results{incorretChar, true, openers}
    }
}

func getIllegalCharacters(lines []string) (string, []string) {
    resultsChan := make(chan results, len(lines))

    for _, line := range lines {
        go findFirstIncorrectChar(line, resultsChan)
    }

    illegalCharacters := ""
    var incompleteLines []string

    for i:=0; i<len(lines); i++ {
        result := <- resultsChan

        if (result.isCorrupt) {
            illegalCharacters += result.char
        } else {
            incompleteLines = append(incompleteLines, result.balance)
        }
    }
    return illegalCharacters, incompleteLines
}

func autoComplete(line string, autoCompleteChan chan string) {
    var balance string

    for i:=len(line)-1; i>=0; i-- {
        balance += getMatch(rune(line[i]))
    }

    autoCompleteChan <- balance
}

func autoCompleteAll(incompleteLines []string) []string {
    autoCompleteChan := make(chan string, len(incompleteLines))

    for _, line := range incompleteLines {
        go autoComplete(line, autoCompleteChan)
    }

    var closers []string
    for i:=0; i<len(incompleteLines); i++ {
        closers = append(closers, <- autoCompleteChan)
    }
    return closers
}

func getPoints(illegalCharacters string) int {
    points := 0

    for _, char := range illegalCharacters {
        switch string(char) {
        case ")":
            points += 3
        case "]":
            points += 57
        case "}":
            points += 1197
        case ">":
            points += 25137
        }
    }
    return points
}

func getAutocompletePoints(closers []string) int {
    var allPoints []int

    for _, closer := range closers {
        points := 0

        for _, char := range closer {
            points *= 5

            switch string(char) {
            case ")":
                points += 1
            case "]":
                points += 2
            case "}":
                points += 3
            case ">":
                points += 4
            }
        }

        allPoints = append(allPoints, points)
    }

    sort.Ints(allPoints)
    return allPoints[(len(allPoints)-1) / 2]
}

func main() {
    lines := ParseFile(10, false)

    // Part 1
    illegalCharacters, incompleteLines := getIllegalCharacters(lines)
    fmt.Printf("Part %d: %d\n", 1, getPoints(illegalCharacters))

    // Part 2
    closers := autoCompleteAll(incompleteLines)
    fmt.Printf("Part %d: %d\n", 2, getAutocompletePoints(closers))
}
