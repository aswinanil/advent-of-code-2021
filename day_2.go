package main

import (
    "fmt"
)

var x, y, aim = 0, 0, 0

func computeDestination(course [][]string) {
    for _, directions := range course {
        distance := GetInt(directions[1])

        switch (directions[0]) {
        case "forward":
            x += distance
        case "up":
            y -= distance
        case "down":
            y += distance
        }
    }
}

func computeDestinationWithAim(course [][]string) {
    for _, directions := range course {
        val := GetInt(directions[1])

        switch (directions[0]) {
        case "forward":
            x += val
            y += aim * val
        case "up":
            aim -= val
        case "down":
            aim += val
        }
    }
}

func main() {
    splitLines := ParseFileAndSplit(2, false)

    // Part 1
    computeDestination(splitLines)
    fmt.Println(x * y)

    // Part 2
    x, y, aim = 0, 0, 0
    computeDestinationWithAim(splitLines)
    fmt.Println(x * y)
}
