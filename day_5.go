package main

import (
    "fmt"
    "strings"
)

func markGrid(line []int, grid [][]int) {
    x1, y1, x2, y2 := line[0], line[1], line[2], line[3]

    if (x1 != x2 && y1 != y2) {
        return
    }

    // Mark horizontally
    xMin, xMax := GetMinMax(x1, x2)
    if (xMin != -1) {
        for x:= xMin; x <= xMax; x++ {
            grid[y1][x]++
        }
        return
    }

    // Mark vertically
    yMin, yMax := GetMinMax(y1, y2)
    if (yMin != -1) {
        for y:= yMin; y <= yMax; y++ {
            grid[y][x1]++
        }
        return
    }
}

func fillGrid(lines [][]int, grid [][]int) {
    for _, line := range lines {
        markGrid(line, grid)
    }
}

func createGrid(lines [][]int) [][]int {
    max := 0

    for _, line := range lines {
        for _, val := range line {
            if val > max {
                max = val
            }
        }
    }

    grid := make([][]int, max+1)

    for i := range grid {
        grid[i] = make([]int, max+1)
    }

    return grid
}

func ParseLines(lines []string) [][]int {
    linesInInt := make([][]int, 0)

    for _, line := range lines {
        coordinates := strings.Split(line, " -> ")

        src := SplitIntoInt(coordinates[0], ",")
        dest := SplitIntoInt(coordinates[1], ",")
        srcAndDest := append(src, dest...)
        linesInInt = append(linesInInt, srcAndDest)
    }

    return linesInInt
}

func getScore(grid [][]int) int {
    score := 0

    for _, line := range grid {
        for _, val := range line {
            if val > 1 {
                score++
            }
        }
    }

    return score
}

func main() {
    lines := ParseFile(5, false)
    linesInInt := ParseLines(lines)

    grid := createGrid(linesInInt)
    // PrintIntArr(grid)

    fillGrid(linesInInt, grid)
    // PrintIntArr(grid)

    fmt.Println(getScore(grid))
}
