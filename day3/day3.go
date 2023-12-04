package main

import (
	"advent-of-code-2023/utils"
	"bufio"
	"fmt"
	"log"
)

type Part struct {
	number int
	row    int
	column int
	size   int
}

var debugRow int = 1

func main() {
	part2()
}

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

	sum := sumParts(lines, 0)
	sum += sumParts(lines, 1)

	for scanner.Scan() {
		lines[0] = lines[1]
		lines[1] = lines[2]
		lines[2] = []rune(scanner.Text())
		sum += sumParts(lines, 1)
	}

	sum += sumParts(lines, 2)

	fmt.Printf("Solution : %v\n", sum)
}

func part2() {
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

	sum := sumGears(lines, 0)
	sum += sumGears(lines, 1)

	for scanner.Scan() {
		lines[0] = lines[1]
		lines[1] = lines[2]
		lines[2] = []rune(scanner.Text())
		sum += sumGears(lines, 1)
	}

	sum += sumGears(lines, 2)

	fmt.Printf("Solution : %v\n", sum)
}

func sumParts(lines [][]rune, row int) int {
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

func sumGears(lines [][]rune, row int) int {
	sum := 0

	fmt.Printf("Row %v :", debugRow)
	debugRow++

	gearPositions := findGearPositions(lines[row])
	for i, gearPosition := range gearPositions {
		numbers := make([]int, 0)
		if row > 0 {
			numbersAbove := searchNumbersAround(lines[row-1], gearPosition)
			numbers = append(numbers, numbersAbove...)
		}

		numbersSides := searchNumbersAround(lines[row], gearPosition)
		numbers = append(numbers, numbersSides...)

		if row < len(lines) {
			numbersBelow := searchNumbersAround(lines[row+1], gearPosition)
			numbers = append(numbers, numbersBelow...)
		}

		if len(numbers) < 2 {
			continue
		}

		if len(numbers) > 2 {
			log.Fatalf("Gear %v of row %v has more than 2 numbers?", i, row)
		}

		sum += numbers[0] * numbers[1]
	}

	fmt.Printf(" Sum: %v\n", sum)

	return sum
}

func findGearPositions(line []rune) []int {
	gearSymbols := make([]int, 0)

	for i, r := range line {
		if r == '*' {
			gearSymbols = append(gearSymbols, i)
		}
	}

	return gearSymbols
}

func searchNumbersAround(line []rune, center int) []int {
	numbers := make([]int, 0)
	start := max(0, center-1)
	end := min(center+1, len(line)-1)

	for i := start; i <= end; i++ {
		if utils.IsDigit(line[i]) {
			number, lastPosition := readNumber(line, i)
			numbers = append(numbers, number)
			i = lastPosition // skip all digits that were read
		}
	}

	return numbers
}

func readNumber(line []rune, startPosition int) (number int, lastPosition int) {
	firstPosition := startPosition
	for ; firstPosition >= 0; firstPosition-- {
		if !utils.IsDigit(line[firstPosition]) {
			break
		}
	}

	currentNumber := 0
	lastPosition = firstPosition + 1
	for ; lastPosition < len(line); lastPosition++ {
		r := line[lastPosition]
		if utils.IsDigit(r) {
			currentNumber = currentNumber*10 + utils.ToInt(r)
		} else {
			break
		}
	}

	return currentNumber, lastPosition - 1
}
