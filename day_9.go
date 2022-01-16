package main

import (
    "fmt"
    "sync"
    "sort"
)

type coordinates struct {
    x int
    y int
}

func getValidAdjacentPoints(x int, y int, heightMap [][]int) []coordinates {
    var validPoints []coordinates

    height := len(heightMap)
    width := len(heightMap[0])

    above := y - 1
    if above >= 0 {
        validPoints = append(validPoints, coordinates{x, above})
    }

    below := y + 1
    if below < height {
        validPoints = append(validPoints, coordinates{x, below})
    }

    left := x - 1
    if left >= 0 {
        validPoints = append(validPoints, coordinates{left, y})
    }

    right := x + 1
    if right < width {
        validPoints = append(validPoints, coordinates{right, y})
    }

    return validPoints
}

func getLowPoints(heightMap [][]int, i int, rowsWg *sync.WaitGroup, lowPointsChan chan coordinates) {
    for j, height := range heightMap[i] {
        adjacentPoints := getValidAdjacentPoints(j, i, heightMap)

        for k, point := range adjacentPoints {
            adjacentHeight := heightMap[point.y][point.x]

            if height >= adjacentHeight {
                break
            }

            if k == len(adjacentPoints) - 1 {
                lowPointsChan <- coordinates{j, i}
            }
        }
    }

    defer rowsWg.Done()
}

func getAllLowPoints(heightMap [][]int) []coordinates {
    lowPointsChan := make(chan coordinates, len(heightMap) * len(heightMap[0]))

    var rowsWg sync.WaitGroup
    rowsWg.Add(len(heightMap))

    for i, _ := range heightMap {
        go getLowPoints(heightMap, i, &rowsWg, lowPointsChan)
    }

    rowsWg.Wait()
    close(lowPointsChan)

    sum := 0
    var lowPoints []coordinates

    for point := range lowPointsChan {
        lowPoints = append(lowPoints, point)
        sum += heightMap[point.y][point.x] + 1
    }

    fmt.Printf("Part %d: %d\n", 1, sum)
    return lowPoints
}

func getAdjacentSlopes(heightMap [][]int, point coordinates) []coordinates {
    pointHeight := heightMap[point.y][point.x]
    var adjacentSlopes []coordinates
    possibleSlopes := getValidAdjacentPoints(point.x, point.y, heightMap)

    for _, slope := range possibleSlopes {
        newHeight := heightMap[slope.y][slope.x]

        if newHeight == 9 {
            continue
        }

        if newHeight <= pointHeight {
            continue
        }

        adjacentSlopes = append(adjacentSlopes, slope)
    }

    return adjacentSlopes
}

func getBasin(heightMap [][]int, lowPoint coordinates, basinsChan chan []coordinates) {
    var newSlopes []coordinates
    processedSlopes := make(map[coordinates]bool)
    newSlopes = append(newSlopes, lowPoint)

    for len(newSlopes) > 0 {
        var newSlope coordinates
        newSlope, newSlopes = newSlopes[0], newSlopes[1:]
        adjacentSlopes := getAdjacentSlopes(heightMap, newSlope)

        var newAdjacentSlopes []coordinates
        for _, adjacentSlope := range adjacentSlopes {
            if processedSlopes[adjacentSlope] {
                continue
            }
            newAdjacentSlopes = append(newAdjacentSlopes, adjacentSlope)
        }

        newSlopes = append(newSlopes, newAdjacentSlopes...)
        processedSlopes[newSlope] = true
    }

    basin := make([]coordinates, len(processedSlopes))

    i := 0
    for key := range processedSlopes {
        basin[i] = key
        i++
    }

    basinsChan <- basin
}

func getAllBasins(heightMap [][]int, lowPoints []coordinates) [][]coordinates {
    basinsChan := make(chan []coordinates, len(lowPoints))

    for _, point := range lowPoints {
        go getBasin(heightMap, point, basinsChan)
    }

    var basins [][]coordinates
    for i:=0; i<len(lowPoints); i++ {
        basins = append(basins, <- basinsChan)
    }

    return basins
}

func getHeightMap(lines []string) ([][]int) {
    heightMap := make([][]int, 0)

    for _, line := range lines {
        var heightRow []int

        for _, char := range line {
            heightRow = append(heightRow, GetInt(string(char)))
        }

        heightMap = append(heightMap, heightRow)
    }

    return heightMap
}

func printLargestBasinsSizeProduct(basins [][]coordinates) {
    var sizes []int

    for _, basin := range basins {
        sizes = append(sizes, len(basin))
    }
    sort.Ints(sizes)

    product := 1
    for i:=len(sizes)-1; i>len(sizes)-4; i-- {
        product *= sizes[i]
    }
    fmt.Printf("Part %d: %d\n", 2, product)
}

func main() {
    lines := ParseFile(9, false)
    heightMap := getHeightMap(lines)

    // Part 1
    lowPoints := getAllLowPoints(heightMap)

    // Part 2
    basins := getAllBasins(heightMap, lowPoints)
    printLargestBasinsSizeProduct(basins)
}
