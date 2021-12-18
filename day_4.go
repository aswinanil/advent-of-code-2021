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

func isLastWinner(winningIndex int, drawSequence int, winnersList []int) bool {
    if drawSequence == 0 {
        return false
    }

    if drawSequence == len(winnersList) - 1 {  // Last draw
        return true
    }

    winnersList[winningIndex] = drawSequence

    for _, val := range winnersList {
        if val == 0 {
            return false
        }
    }

    return true
}

func markBoards(draws []int, allBoards [][][]int, isGetLast bool) (int, int) {
    winnersList := make([]int, len(allBoards))

    for i, draw := range draws {
        for j, board := range allBoards {
            isWinner := markBoard(draw, board)

            if (isWinner && (!isGetLast || isLastWinner(j, i, winnersList))) {
                return j, draw
            }
        }
    }

    return -1, -1
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
    allBoardsCopy := make([][][]int, len(allBoards))
    copy(allBoardsCopy, allBoards)
    winningIndex, lastDraw := markBoards(draws, allBoardsCopy, false)
    fmt.Println(getScore(allBoardsCopy[winningIndex], lastDraw))

    // Part 2
    winningIndex, lastDraw = markBoards(draws, allBoards, true)
    fmt.Println(getScore(allBoards[winningIndex], lastDraw))
}
