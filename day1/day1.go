package main

import (
	"advent-of-code-2023/utils"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var digitWords [9]string = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	part2()
}

func part1() {
	file, err := os.Open("day1-input.txt")
	utils.Check(err)
	defer file.Close()

	sum := 0
	reader := bufio.NewReader(file)
	for {
		calibrationValue, err := readCalibrationValue(reader)
		sum += calibrationValue

		if err == io.EOF {
			break
		}

	}

	fmt.Printf("Sum: %v", sum)
}

func part2() {
	file, err := os.Open("day1-input.txt")
	utils.Check(err)
	defer file.Close()

	sum := 0
	reader := bufio.NewReader(file)
	i := 1
	for {
		calibrationValue, err := readCalibrationValuePart2(reader)

		fmt.Printf("%v - Calibration value : %v\n", i, calibrationValue)
		i++

		sum += calibrationValue

		if err == io.EOF {
			break
		}
	}

	fmt.Printf("Sum: %v", sum)
}

func readCalibrationValue(reader *bufio.Reader) (calibrationValue int, err error) {
	// fmt.Println("New line")

	firstDigit := -1
	lastDigit := -1

	var readError error = nil
	for {
		r, _, err := reader.ReadRune()
		readError = err

		if r == '\r' || err == io.EOF { // running on windows w/e...
			break
		}

		if utils.IsDigit(r) {
			i := utils.ToInt(r)
			if firstDigit == -1 {
				firstDigit = i
			}
			lastDigit = i
		}
	}

	// fmt.Printf("First digit : %v\n", firstDigit)
	// fmt.Printf("Last digit : %v\n", lastDigit)

	s := fmt.Sprintf("%v%v", firstDigit, lastDigit)
	i, err := strconv.Atoi(s)
	utils.Check(err)

	// fmt.Printf("Calibration value : %v\n", i)

	return i, readError
}

func readCalibrationValuePart2(reader *bufio.Reader) (calibrationValue int, err error) {
	firstDigit := -1
	lastDigit := -1
	currentWord := ""

	var updateDigits = func(digit int) {
		if firstDigit == -1 {
			firstDigit = digit
		}
		lastDigit = digit
	}

	var seekBack = func(offset int) {
		wordLength := len(currentWord)
		currentWord = currentWord[wordLength-offset : wordLength]
	}

	var isSpelledDigit = func() bool {
		isMatch := false
		for i, digitWord := range digitWords {
			if digitWord == currentWord {
				updateDigits(i + 1)
				seekBack(1) // keep only last character
				isMatch = true
				break
			} else if strings.HasPrefix(digitWord, currentWord) {
				isMatch = true
				break
			}
		}

		return isMatch
	}

	var readError error = nil
	for {
		r, _, err := reader.ReadRune()
		readError = err

		if r == '\r' || err == io.EOF { // running on windows w/e...
			break
		}

		currentWord = currentWord + string(r)
		fmt.Printf("%v - Word : %v\n", string(r), currentWord)

		if utils.IsDigit(r) {
			i := utils.ToInt(r)
			updateDigits(i)
			currentWord = ""
			continue
		}

		if !isSpelledDigit() {
			seekBack(len(currentWord) - 1) // remove only first character
		}
	}

	// fmt.Printf("First digit : %v\n", firstDigit)
	// fmt.Printf("Last digit : %v\n", lastDigit)

	i := firstDigit*10 + lastDigit

	return i, readError
}
