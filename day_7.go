package main

import (
    "fmt"
    "sync"
)

func getExpensiveFuelCost(targetX int, x int) int {
    diff := Abs(x - targetX)
    fuel := diff

    for i:=1; i<diff; i++ {
        fuel += i
    }

    return fuel
}

func getFuel(positions []int, targetX int, isExpensive bool, fuelChan chan int) {
    fuel := 0

    for _, x := range positions {
        if (isExpensive) {
            fuel += getExpensiveFuelCost(targetX, x)
        } else {
            fuel += Abs(x - targetX)
        }
    }

    fuelChan <- fuel
}

func getMinFuel(positions []int, isExpensive bool, wg *sync.WaitGroup, part int) {
    max := GetMax(positions)
    fuelChan := make(chan int, max)

    for i:=0; i<max; i++ {
        go getFuel(positions, i, isExpensive, fuelChan)
    }

    minFuel := <- fuelChan

    for i:=0; i < max - 1; i++ {
        fuelCost := <- fuelChan

        if fuelCost < minFuel {
            minFuel = fuelCost
        }
    }

    fmt.Printf("Part %d: %d\n", part, minFuel)
    defer wg.Done()
}

func main() {
    lines := ParseFile(7, false)
    positions := SplitIntoInt(lines[0], ",")

    var wg sync.WaitGroup
    wg.Add(2)

    go getMinFuel(positions, false, &wg, 1)  // Part 1
    go getMinFuel(positions, true, &wg, 2)  // Part 2

    wg.Wait()
}
