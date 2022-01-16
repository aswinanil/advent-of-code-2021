package main

import (
    "fmt"
)

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

func findFirstIncorrectChar(line string, incorretCharChan chan string) {
    var openers string
    var incorretChar string

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

    incorretCharChan <- incorretChar
}

func getIllegalCharacters(lines []string) string {
    incorretCharChan := make(chan string, len(lines))

    for _, line := range lines {
        go findFirstIncorrectChar(line, incorretCharChan)
    }

    illegalCharacters := ""
    for i:=0; i<len(lines); i++ {
        illegalCharacters += <- incorretCharChan
    }
    return illegalCharacters
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

func main() {
    lines := ParseFile(10, false)

    illegalCharacters := getIllegalCharacters(lines)
    fmt.Printf("Part %d: %d\n", 1, getPoints(illegalCharacters))
}
