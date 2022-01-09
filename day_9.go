package main

import (
    "fmt"
    "sync"
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

func getRiskLevelsSum(heightMap [][]int, wg *sync.WaitGroup, part int) {
    lowPointsChan := make(chan coordinates, len(heightMap) * len(heightMap[0]))

    var rowsWg sync.WaitGroup
    rowsWg.Add(len(heightMap))

    for i, _ := range heightMap {
        go getLowPoints(heightMap, i, &rowsWg, lowPointsChan)
    }

    rowsWg.Wait()
    close(lowPointsChan)

    sum := 0
    for point := range lowPointsChan {
        sum += heightMap[point.y][point.x] + 1
    }

    fmt.Printf("Part %d: %d\n", part, sum)
    defer wg.Done()
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

func main() {
    lines := ParseFile(9, false)

    heightMap := getHeightMap(lines)

    var wg sync.WaitGroup
    wg.Add(1)

    go getRiskLevelsSum(heightMap, &wg, 1)

    wg.Wait()
}
