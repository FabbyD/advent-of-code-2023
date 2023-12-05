package main

import (
	"advent-of-code-2023/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers map[int]struct{}
	numbers        []int
}

func (card *Card) IsWinningNumber(number int) bool {
	_, ok := card.winningNumbers[number]
	return ok
}

func main() {
	part2()
}

func part1() {
	file := utils.OpenInput()
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += scoreCard(scanner.Text())
	}

	fmt.Printf("Solution : %v\n", sum)
}

func part2() {
	file := utils.OpenInput()
	defer file.Close()

	scanner := bufio.NewScanner(file)

	copies := make(map[int]int)

	numCards := 0
	for scanner.Scan() {
		copies = findCardCopies(scanner.Text(), copies)
		numCards++
	}

	solution := numCards
	for i := 0; i < numCards; i++ {
		solution += copies[i+1]
	}

	fmt.Printf("Solution : %v\n", solution)
}

func scoreCard(cardLine string) int {
	score := 0

	card := parseCard(cardLine)

	fmt.Println(card)

	for _, number := range card.numbers {
		if !card.IsWinningNumber(number) {
			continue
		}

		if score == 0 {
			score = 1
		} else {
			score *= 2
		}
	}

	fmt.Printf("%v => Score : %v\n", cardLine, score)

	return score
}

func findCardCopies(cardLine string, copies map[int]int) map[int]int {
	card := parseCard(cardLine)

	numWinningNumbers := 0
	for _, number := range card.numbers {
		if card.IsWinningNumber(number) {
			numWinningNumbers++
		}
	}

	start := card.id + 1
	for i := start; i < start+numWinningNumbers; i++ {
		copies[i] += copies[card.id] + 1
	}

	return copies
}

func parseCard(cardLine string) Card {
	tokens := strings.Split(cardLine, ":")

	cardPart := tokens[0]
	numbersPart := tokens[1]

	tokens = strings.Split(cardPart, " ")
	cardId, err := strconv.Atoi(tokens[len(tokens)-1])
	utils.Check(err)

	tokens = strings.Split(numbersPart, "|")
	winningNumbersPart, numbersPart := tokens[0], tokens[1]

	return Card{
		id:             cardId,
		winningNumbers: makeSet(splitNumbers(winningNumbersPart)),
		numbers:        splitNumbers(numbersPart),
	}
}

func splitNumbers(numbersString string) []int {
	numbersString = strings.TrimSpace(numbersString)

	numbers := make([]int, 0)
	for _, numberString := range strings.Split(numbersString, " ") {
		n, err := strconv.Atoi(numberString)
		if err != nil {
			continue
		}

		numbers = append(numbers, n)
	}

	return numbers
}

func makeSet(numbers []int) map[int]struct{} {
	set := make(map[int]struct{})

	for _, number := range numbers {
		set[number] = struct{}{}
	}

	return set
}
