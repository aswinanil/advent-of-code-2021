package main

import (
    "fmt"
)

var x, y = 0, 0

func computeDestination(course [][]string) {
    for _, direction := range course {
        distance := GetInt(direction[1])

        switch (direction[0]) {
        case "forward":
            x += distance
        case "up":
            y -= distance
        case "down":
            y += distance
        }
    }
}

func main() {
    splitLines := ParseFileAndSplit(2, false)
    computeDestination(splitLines)
    fmt.Println(x * y)
}
