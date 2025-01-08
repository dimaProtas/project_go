package main

import (
	"errors"
	"fmt"
	"os"
)

const row, col = 9, 9

type sudoku [row][col]int8

var (
	ErrRow          = errors.New("row out of range")
	ErrCol          = errors.New("column out of range")
	ErrNum          = errors.New("number out of range")
	ErrInsert       = errors.New("нельзя вставить число в эту ячейку та как она фиксированная!")
	ErrRowInser     = errors.New("в этом ряду уже есть такое сичло!")
	ErrColInser     = errors.New("в этом столбце уже есть такое сичло!")
	ErrRegionInsert = errors.New("В этой группе уже есть такое число!")
)

// printSudoku печатает доску Судоку с разметкой
func printSudoku(s sudoku) {
	fmt.Println("╔═══════╦═══════╦═══════╗")
	for i, row := range s {
		fmt.Print("║ ")
		for j, cell := range row {
			if cell == 0 {
				fmt.Print(". ") // Пустые клетки обозначаем точками
			} else {
				fmt.Printf("%d ", cell)
			}
			if (j+1)%3 == 0 && j != 8 {
				fmt.Print("║ ")
			}
		}
		fmt.Println("║")
		if (i+1)%3 == 0 && i != 8 {
			fmt.Println("╠═══════╬═══════╬═══════╣")
		}
	}
	fmt.Println("╚═══════╩═══════╩═══════╝")
}

func (s *sudoku) clearKvadrat(n int8) error {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			switch n {
			case 1:
				if r < 3 && c < 3 {
					s[r][c] = 0
				}
			case 2:
				if (c > 2 && c < 6) && r < 3 {
					s[r][c] = 0
				}
			case 3:
				if (c > 5 && c < 9) && r < 3 {
					s[r][c] = 0
				}
			case 4:
				if (c < 3) && (r > 2 && r < 6) {
					s[r][c] = 0
				}
			case 5:
				if (c > 2 && c < 6) && (r > 2 && r < 6) {
					s[r][c] = 0
				}
			case 6:
				if (c > 5 && c < 9) && (r > 2 && r < 6) {
					s[r][c] = 0
				}
			case 7:
				if (c < 3) && (r > 5 && r < 9) {
					s[r][c] = 0
				}
			case 8:
				if (c > 2 && c < 6) && (r > 5 && r < 9) {
					s[r][c] = 0
				}
			case 9:
				if (c > 5 && c < 9) && (r > 5 && r < 9) {
					s[r][c] = 0
				}
			}
		}
	}

	return nil
}

func (s sudoku) isRow(r, n int8) bool {
	for c := 0; c < col; c++ {
		if s[r-1][c] == n {
			return true
		}
	}
	return false
}

func (s sudoku) isCol(c, n int8) bool {
	for r := 0; r < row; r++ {
		if s[r][c-1] == n {
			return true
		}
	}
	return false
}

func (s sudoku) inRegion(row, column int8, digit int8) bool {
	startRow, startColumn := row/3*3, column/3*3
	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn+3; c++ {
			if s[r][c] == digit {
				return true
			}
		}
	}
	return false
}

func (s *sudoku) insertNum(r, c, n int8) error {
	if r < 1 || r > 9 {
		return ErrRow
	}
	if c < 1 || c > 9 {
		return ErrCol
	}
	if n < 1 || n > 9 {
		return ErrNum
	}
	if s[r-1][c-1] != 0 {
		return ErrInsert
	}

	if s.isRow(r, n) {
		return ErrRowInser
	}

	if s.isCol(c, n) {
		return ErrColInser
	}
	if s.inRegion(r-1, c-1, n) {
		return ErrRegionInsert
	}
	s[r-1][c-1] = n
	return nil
}

func main() {
	sudokuGame := sudoku{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9}}

	err := sudokuGame.insertNum(1, 1, 2)
	if err != nil {
		fmt.Printf("ОШибка: %v\n", err)
		os.Exit(1)
	}

	// sudokuGame.clearKvadrat(9)

	printSudoku(sudokuGame)
}
