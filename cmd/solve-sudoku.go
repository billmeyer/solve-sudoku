package main

import (
	"fmt"
	"math"
	"time"
)

// Programming for the Puzzled -- Srini Devadas
// You Will Never Want to Play Sudoku Again
// Given a partially filled in Sudoku board, complete the puzzle
// obeying the rules of Sudoku

// Global variable set to 0
var backtracks = 0

var input = [9][9]int{
	{5, 1, 7, 6, 0, 0, 0, 3, 4},
	{2, 8, 9, 0, 0, 4, 0, 0, 0},
	{3, 4, 6, 2, 0, 5, 0, 9, 0},
	{6, 0, 2, 0, 0, 0, 0, 1, 0},
	{0, 3, 8, 0, 0, 6, 0, 4, 7},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 9, 0, 0, 0, 0, 0, 7, 8},
	{7, 0, 3, 4, 0, 0, 5, 6, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

var inp2 = [9][9]int{
	{5, 1, 7, 6, 0, 0, 0, 3, 4},
	{0, 8, 9, 0, 0, 4, 0, 0, 0},
	{3, 0, 6, 2, 0, 5, 0, 9, 0},
	{6, 0, 0, 0, 0, 0, 0, 1, 0},
	{0, 3, 0, 0, 0, 6, 0, 4, 7},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 9, 0, 0, 0, 0, 0, 7, 8},
	{7, 0, 3, 4, 0, 0, 5, 6, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

var inpd = [9][9]int{
	{1, 0, 5, 7, 0, 2, 6, 3, 8},
	{2, 0, 0, 0, 0, 6, 0, 0, 5},
	{0, 6, 3, 8, 4, 0, 2, 1, 0},
	{0, 5, 9, 2, 0, 1, 3, 8, 0},
	{0, 0, 2, 0, 5, 8, 0, 0, 9},
	{7, 1, 0, 0, 3, 0, 5, 0, 2},
	{0, 0, 4, 5, 6, 0, 7, 2, 0},
	{5, 0, 0, 0, 0, 4, 0, 6, 3},
	{3, 2, 6, 1, 0, 7, 0, 0, 4},
}

var hard = [9][9]int{
	{8, 5, 0, 0, 0, 2, 4, 0, 0},
	{7, 2, 0, 0, 0, 0, 0, 0, 9},
	{0, 0, 4, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 7, 0, 0, 2},
	{3, 0, 5, 0, 0, 0, 9, 0, 0},
	{0, 4, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 8, 0, 0, 7, 0},
	{0, 1, 7, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 3, 6, 0, 4, 0},
}

var diff = [9][9]int{
	{0, 0, 5, 3, 0, 0, 0, 0, 0},
	{8, 0, 0, 0, 0, 0, 0, 2, 0},
	{0, 7, 0, 0, 1, 0, 5, 0, 0},
	{4, 0, 0, 0, 0, 5, 3, 0, 0},
	{0, 1, 0, 0, 7, 0, 0, 0, 6},
	{0, 0, 3, 2, 0, 0, 0, 8, 0},
	{0, 6, 0, 5, 0, 0, 0, 0, 9},
	{0, 0, 4, 0, 0, 0, 0, 3, 0},
	{0, 0, 0, 0, 0, 9, 7, 0, 0},
}

var expert = [9][9]int{
	{0, 0, 6, 7, 8, 0, 0, 0, 0},
	{0, 8, 5, 0, 0, 4, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 8, 0, 9},
	{6, 0, 0, 3, 0, 2, 4, 0, 0},
	{3, 5, 0, 0, 1, 0, 0, 2, 7},
	{0, 0, 1, 6, 0, 7, 0, 0, 5},
	{9, 0, 2, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 4, 0, 0, 2, 1, 0},
	{0, 0, 0, 0, 7, 1, 9, 0, 0},
}

// This procedure finds the next empty square to fill on the Sudoku grid
func findNextCellToFill(grid *[9][9]int) (int, int) {
	// Look for an unfilled grid location
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if grid[x][y] == 0 {
				return x, y
			}
		}
	}
	return -1, -1
}

// This procedure checks if setting the (i, j) square to e is valid
func isValid(grid *[9][9]int, i, j, e int) bool {
	rowOk := true
	//rowOk = all([e != grid[i][x] for x in range(9)])
	for x := 0; x < 9; x++ {
		if e == grid[i][x] {
			rowOk = false
			break
		}
	}

	if rowOk {
		columnOk := true
		//columnOk = all([e != grid[x][j] for x in range (9)])
		for x := 0; x < 9; x++ {
			if e == grid[x][j] {
				columnOk = false
				break
			}
		}

		if columnOk {
			// finding the top left x,y co-ordinates of
			// the section or sub-grid containing the i,j cell
			secTopX := 3 * int(math.Floor(float64(i)/3))
			secTopY := 3 * int(math.Floor(float64(j)/3))
			for x := secTopX; x < secTopX+3; x++ {
				for y := secTopY; y < secTopY+3; y++ {
					if grid[x][y] == e {
						return false
					}
				}
			}
			return true
		}
	}
	return false
}

// This procedure fills in the missing squares of a Sudoku puzzle
// obeying the Sudoku rules through brute-force guessing and checking
func solveSudoku(grid *[9][9]int, i, j int) bool {

	// find the next cell to fill
	i, j = findNextCellToFill(grid)
	if i == -1 {
		return true
	}

	for e := 1; e < 10; e++ {
		// Try different values in i, j location
		if isValid(grid, i, j, e) {
			grid[i][j] = e
			if solveSudoku(grid, i, j) {
				return true
			}

			// Undo the current cell for backtracking
			backtracks += 1
			grid[i][j] = 0
		}
	}

	return false
}

func printSudoku(grid [9][9]int) {
	numrow := 0
	for r := 0; r < len(grid); r++ {
		row := grid[r]
		if numrow%3 == 0 && numrow != 0 {
			fmt.Println()
		}
		fmt.Printf("%v %v %v\n", row[0:3], row[3:6], row[6:9])
		numrow += 1
	}
	return
}

func main() {
	//fmt.Println(solveSudoku(&input, 0, 0))
	//printSudoku(input)
	//fmt.Printf("Backtracks = %v\n", backtracks)
	//
	//backtracks = 0
	//fmt.Println(solveSudoku(&inp2, 0, 0))
	//printSudoku(inp2)
	//fmt.Printf("Backtracks = %v\n", backtracks)
	//
	//backtracks = 0
	//printSudoku(hard)
	//fmt.Println(solveSudoku(&hard, 0, 0))
	//printSudoku(hard)
	//fmt.Printf("Backtracks = %v\n", backtracks)
	//
	//backtracks = 0
	//printSudoku(diff)
	//fmt.Println(solveSudoku(&diff, 0, 0))
	//printSudoku(hard)
	//fmt.Printf("Backtracks = %v\n", backtracks)

	backtracks = 0
	fmt.Printf("Starting Puzzle\n")
	fmt.Printf("===============\n")
	printSudoku(expert)

	startTime := time.Now()
	solved := solveSudoku(&expert, 0, 0)
	endTime := time.Now()

	if solved == false {
		fmt.Printf("No solution found\n")
	} else {
		fmt.Printf("\nSolution\n")
		fmt.Printf("===============\n")
		printSudoku(expert)
		fmt.Printf("Backtracks = %v\n", backtracks)
	}

	fmt.Printf("cpu time: %v\n", endTime.Sub(startTime))
}
