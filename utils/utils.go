package utils

import (
	"log"
	"os"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func OpenInput() *os.File {
	file, err := os.Open("input.txt")
	Check(err)

	return file
}

func OpenExampleInput(day int) *os.File {
	file, err := os.Open("example.txt")
	Check(err)

	return file
}

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func ToInt(r rune) int {
	return int(r - '0')
}
