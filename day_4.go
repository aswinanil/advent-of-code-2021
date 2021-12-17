package main

import (
    "fmt"
)

var boardLength int
var boardHeight int

type mark struct {
    x int
    y int
}

type results struct {
    a int
    b int
}

func parseBoard(lines []string) ([]int, [][][]int) {
    draws := SplitIntoInt(lines[0], ",")

    var board [][]int
    var allBoards [][][]int

    for i:= 2; i < len(lines); i++ {
        line := lines[i]

        if line == "" {
            allBoards = append(allBoards, board)
            board = nil
            continue
        }

        board = append(board, SplitIntoInt(line, " "))

        if i == len(lines) - 1 {
            allBoards = append(allBoards, board)
        }
    }

    return draws, allBoards
}

func checkBoard(board [][]int, x int, y int) bool {
    row := board[y]
    for i:=0; i<boardLength; i++ {
        if (row[i] != -1) {
            break
        }
        if i == boardLength - 1 {
            return true
        }
    }

    for i:=0; i<boardHeight; i++ {
        if (board[i][x] != -1) {
            break
        }
        if i == boardHeight - 1 {
            return true
        }
    }

    return false
}

func markBoard(draw int, board [][]int) bool {
    for y:= 0; y < len(board); y++ {
        for x:= 0; x < len(board[0]); x++ {
            if board[y][x] == draw {
                board[y][x] = -1
                if checkBoard(board, x, y) {
                    return true
                }
            }
        }
    }
    return false
}

func markBoards(draws []int, allBoards[][][]int) (int, int) {
    for _, draw := range draws {
        for j, board := range allBoards {
            if markBoard(draw, board) {
                return j, draw
            }
        }
    }

    return -1, -1
}

func printBoard(board [][]int) {
    for _, row := range board {
        fmt.Println(row)
    }
    fmt.Println("")
}

func getScore(board [][]int, lastDraw int) int {
    score := 0

    for y:=0; y<boardHeight; y++ {
        for x:=0; x<boardLength; x++ {
            if board[y][x] != -1 {
                score += board[y][x]
            }
        }
    }

    return score * lastDraw
}

func main() {
    lines := ParseFile(4, false)
    draws, allBoards := parseBoard(lines)
    boardLength = len(allBoards[0][0])
    boardHeight = len(allBoards[0])

    // Part 1
    winningIndex, lastDraw := markBoards(draws, allBoards)
    fmt.Println(getScore(allBoards[winningIndex], lastDraw))
}
