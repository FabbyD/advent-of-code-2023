package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	part1()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func part1() {
	file, err := os.Open("day1-input.txt")
	check(err)
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

func readCalibrationValue(reader *bufio.Reader) (calibrationValue int, err error) {
	fmt.Println("New line")

	firstDigit := -1
	lastDigit := -1

	var readError error = nil
	for {
		r, _, err := reader.ReadRune()
		readError = err

		if r == '\r' || err == io.EOF { // running on windows w/e...
			break
		}

		fmt.Println(string(r))

		if isDigit(r) {
			i := toInt(r)
			if firstDigit == -1 {
				firstDigit = i
			}
			lastDigit = i
		}
	}

	fmt.Printf("First digit : %v\n", firstDigit)
	fmt.Printf("Last digit : %v\n", lastDigit)

	s := fmt.Sprintf("%v%v", firstDigit, lastDigit)
	i, err := strconv.Atoi(s)
	check(err)

	fmt.Printf("Calibration value : %v\n", i)

	return i, readError
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func toInt(r rune) int {
	return int(r - '0')
}
