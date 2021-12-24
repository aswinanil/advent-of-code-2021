package main

import (
    "fmt"
    "strings"
)

func isDiagonal(line []int) bool {
    x1, y1, x2, y2 := line[0], line[1], line[2], line[3]
    return x1 != x2 && y1 != y2
}

func markGridDiagonally(line []int, grid [][]int) {
    x1, y1, x2, y2 := line[0], line[1], line[2], line[3]

    xMin, xMax := GetMinMax(x1, x2)
    yOfXMin, yOfXMax := -1, -1

    if (xMin == x1) {
        yOfXMin = y1
        yOfXMax = y2
    } else {
        yOfXMin = y2
        yOfXMax = y1
    }

    y := yOfXMin
    isIncY := yOfXMin < yOfXMax

    for x:= xMin; x <= xMax; x++ {
        grid[y][x]++

        if (isIncY) {
            y++
        } else {
            y--
        }
    }

    return
}

func markGrid(line []int, grid [][]int) {
    x1, y1, x2, y2 := line[0], line[1], line[2], line[3]

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

func fillGrid(lines [][]int, grid [][]int, checkDiagonal bool) {
    for _, line := range lines {
        if (!checkDiagonal || !isDiagonal(line)) {
            markGrid(line, grid)
        } else {
            markGridDiagonally(line, grid)
        }
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

    // Part 1
    grid := createGrid(linesInInt)
    fillGrid(linesInInt, grid, false)
    fmt.Println(getScore(grid))

    // Part 2
    gridWithDiag := createGrid(linesInInt)
    fillGrid(linesInInt, gridWithDiag, true)
    fmt.Println(getScore(gridWithDiag))
}
