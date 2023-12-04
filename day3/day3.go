package main

import (
	"advent-of-code-2023/utils"
	"bufio"
	"fmt"
)

type Part struct {
	number int
	row    int
	column int
	size   int
}

func main() {
	part1()
}

var debugRow int = 1

func part1() {
	file := utils.OpenInput()
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]rune, 3)

	scanner.Scan()
	lines[0] = []rune(scanner.Text())
	scanner.Scan()
	lines[1] = []rune(scanner.Text())
	scanner.Scan()
	lines[2] = []rune(scanner.Text())

	sum := processRow(lines, 0)
	sum += processRow(lines, 1)

	for scanner.Scan() {
		lines[0] = lines[1]
		lines[1] = lines[2]
		lines[2] = []rune(scanner.Text())
		sum += processRow(lines, 1)
	}

	sum += processRow(lines, 2)

	fmt.Printf("Solution : %v\n", sum)
}

func processRow(lines [][]rune, row int) int {
	sum := 0

	fmt.Printf("Row %v :", debugRow)
	debugRow++

	numbers := findNumbers(lines[row], row)
	for _, number := range numbers {
		if isPart(lines, number) {
			fmt.Printf(" %v", number.number)
			sum += number.number
		}
	}

	fmt.Printf(" Sum: %v\n", sum)

	return sum
}

func findNumbers(line []rune, row int) []Part {
	parts := make([]Part, 0)

	currentNumber := 0
	column := 0
	for i, r := range line {
		if utils.IsDigit(r) {
			if currentNumber == 0 {
				column = i
			}
			currentNumber = currentNumber*10 + utils.ToInt(r)
		} else if currentNumber > 0 {
			parts = append(parts, Part{
				number: currentNumber,
				row:    row,
				column: column,
				size:   i - column,
			})
			currentNumber = 0
		}
	}

	if currentNumber > 0 {
		parts = append(parts, Part{
			number: currentNumber,
			row:    row,
			column: column,
			size:   len(line) - column,
		})
	}

	return parts
}

func isPart(lines [][]rune, part Part) bool {
	startRow := max(0, part.row-1)
	endRow := min(part.row+1, len(lines)-1)
	startColumn := max(0, part.column-1)
	endColumn := min(part.column+part.size, len(lines[0])-1)

	for i := startRow; i <= endRow; i++ {
		for j := startColumn; j <= endColumn; j++ {
			if isSymbol(lines[i][j]) {
				return true
			}
		}
	}

	return false
}

func isSymbol(r rune) bool {
	return !utils.IsDigit(r) && r != '.' && r != '\n'
}
