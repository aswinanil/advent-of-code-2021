package main

import (
    "fmt"
)

func getFuel(positions []int, i int, fuelChan chan int) {
    fuel := 0
    targetX := positions[i]

    for _, x := range positions {
        fuel += Abs(x - targetX)
    }

    fuelChan <- fuel
}

func getMinFuel(positions []int) int {
    fuelChan := make(chan int, len(positions))

    for i, _ := range positions {
        go getFuel(positions, i, fuelChan)
    }

    minFuel := <- fuelChan

    for i:=0; i < len(positions) - 1; i++ {
        fuelCost := <- fuelChan

        if fuelCost < minFuel {
            minFuel = fuelCost
        }
    }

    return minFuel
}

func main() {
    lines := ParseFile(7, false)
    positions := SplitIntoInt(lines[0], ",")

    // Part 1
    fmt.Println(getMinFuel(positions))
}
