# tetris-optimizer

A Go program that reads a text file containing Tetromino definitions and assembles them into the smallest possible square, respecting order and displaying each Tetromino with uppercase letters (A, B, C, …).

The program also validates input format and prints ERROR on invalid files.

# Features

Assemble tetrominoes into the smallest square possible.

Identify each tetromino with unique uppercase letters.

Handles incomplete squares by leaving spaces.

Validates file and tetromino formats; prints ERROR on invalid input.

Written entirely in Go using standard packages.



Only standard library packages are used (os, fmt, strings, etc.).

Installation

Clone the repository:

git clone https://learn.zone01kisumu.ke/git/piregi/tetris-optimizer



Run the program using Go:

go run . <path_to_tetromino_file>

# Input File Format

    Each Tetromino is represented in a 4x4 grid.

    Empty spaces: .

    Blocks: #

    Tetrominoes are separated by empty lines.

    At least one Tetromino must be present.

Example sample.txt:

...#                                               
...#                                        
...#                                   
...#                                                 
                                                         
....                                                         
....                                                            
..##                                                         
..##                                                           
                                                             
.###                                                            
...#                                                         
....                                                            
....                                                             

# Output

    Prints the square with tetrominoes labeled as A, B, C, etc.

    Spaces (.) are used if a square cannot be completely filled.

    Example output:

ABBBB.                                                           
ACCCEE                                                          
AFFCEE                                                            
A.FFGG                                                         
HHHDDG                                                             
.HDD.G                                                    

    Errors during parsing or invalid file format print:

ERROR

# Usage Example

go run . sample.txt

Output:

ABBBB.                            
ACCCEE                       
AFFCEE                    
A.FFGG                                   
HHHDDG                                     
.HDD.G                             

# Project Structure

tetris-optimizer/         
│            
├── go.mod                   # Go module file                 
├── main.go                  # Entry point                 
├── tetris/                            
│   ├── model/                         
│   │    └── tetromino.go    # Shared structs: Tetromino, Point  
│   ├── Parsing/                            
│   │    └── parsing.go      # Functions to parse text files into Tetromino slices       
│   └── Board/            
│        └── board.go        # Board creation, Solve, PrintBoard
├── README.md       
└── sample.txt               # Example input      

# Functions Overview
Parsing package

    ParseTetrominoes(lines []string) ([]model.Tetromino, bool) — parses tetrominoes from file lines.

    Validates tetromino format.

Board package

    NewBoard(size int) [][]rune — creates a square board.

    Solve(board [][]rune, pieces []model.Tetromino, index int) bool — recursively solves the square.

    PrintBoard(board [][]rune) — prints the board with letters.

# Error Handling

    If the program receives invalid arguments, invalid file format, or malformed tetrominoes, it prints:

ERROR

    The program exits immediately after printing the error.