package main
 
import (
	"os"
	"fmt"
	//"math"
	"strings"
	Board "tetris/tetris/Board"
	parsing "tetris/tetris/Parsing"
	

)
type Point struct {
	x int
	y int
}

type Tetromino struct {
	blocks []Point
	letter rune
}



func main(){
	

	if !ValidArgs(os.Args){
		fmt.Println("ERROR")
		return
	}else{
		filename:=os.Args[1]
		if !ValidFileName(filename){
			fmt.Println("ERROR")
			return
		}else{
			//read file 
			data, err := os.ReadFile(filename)
			if err != nil{
				fmt.Println("ERROR")
			}

			lines := strings.Split(string(data), "\n")


			pieces, ok := parsing.ParseTetrominoes(lines)
			if !ok || len(pieces) == 0{
				fmt.Println("ERROR")
				return
			}

			//start from smallest square

			size := Board.MinBoardSize(len(pieces))

			for {
				board := Board.NewBoard(size)

				if Board.Solve(board, pieces, 0){
					Board.PrintBoard(board)
					return
				}
				size++
			}
			

		}
		

	}
	
}
func ValidArgs(args []string) bool{
	result := false
	if len(args) == 2{
		result = true

	}else {
		result= false
	}
	return result
}

func ValidFileName(fileName string) bool{
	return strings.HasSuffix(fileName, ".txt")
}