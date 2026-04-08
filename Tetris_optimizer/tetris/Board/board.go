package tetris

import (
	"math"
	"fmt"
	model "tetris/tetris/model"
)


func NewBoard(size int) [][]rune{
	board := make ([][]rune, size)

	for i := range board {
		board[i] = make([]rune, size)
		for j := range board {
			board[i][j] = '.'
		}
	}
	return board
}


func CanPlace( board [][]rune, t model.Tetromino, x, y int) bool{
	size := len(board)

	for _, b := range t.Blocks {

		nx := x + b.X
		ny := y + b.Y

		if nx < 0 || ny < 0 || nx >= size || ny >= size {
			return false
		}

		if board[ny][nx] != '.'{
			return false
		}
	}
	return true
}


func Place(board [][]rune, t model.Tetromino, x, y int){
	for _, b := range t.Blocks{
		board[y+b.Y][x+b.X] = t.Letter
	}
}

func Remove(board [][]rune, t model.Tetromino, x, y int){
	for _, b := range t.Blocks {
		board[y+b.Y][x+b.X] = '.'
	}
}

//solver

func Solve(board [][]rune, pieces []model.Tetromino, index int)bool{
	if index == len(pieces){
		return true
	}

	size := len(board)
	current := pieces[index]

	for y := 0; y < size; y++{
		for x := 0; x < size; x++{
			
			if CanPlace(board, current, x, y){
				Place(board, current, x, y)

				if Solve(board, pieces, index+1){
					return true
				}

				Remove(board, current, x, y)
			}
		}
	}

	return false
}

func MinBoardSize(pieceCount int)int {
	total := pieceCount * 4
	return int(math.Ceil(math.Sqrt(float64(total))))
}

func PrintBoard(board [][]rune){
	for _, row := range board {
		fmt.Println(string(row))
	}
}