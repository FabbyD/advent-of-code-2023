package main

import (
	"advent-of-code-2023/utils"
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type HandType int

const (
	FiveOfAKind HandType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

type Card struct {
	card  int
	count int
}

type Play struct {
	hand      string
	cards     []int
	bid       int
	handType  HandType
	cardCount []Card
}

func main() {
	part1()
}

func part1() {
	file := utils.OpenInput()

	scanner := bufio.NewScanner(file)

	plays := make([]Play, 0)
	for scanner.Scan() {
		line := scanner.Text()
		plays = append(plays, parsePlay(line))
	}

	plays = rankPlays(plays)

	sum := 0
	for i, play := range plays {
		fmt.Println(play)
		sum += (i + 1) * play.bid
	}

	fmt.Println("Solution:", sum)
}

func rankPlays(plays []Play) []Play {
	slices.SortFunc(plays, func(lhs, rhs Play) int {
		if lhs.handType > rhs.handType {
			return -1
		}

		if lhs.handType < rhs.handType {
			return 1
		}

		for i := 0; i < len(lhs.hand); i++ {
			leftCard, rightCard := lhs.cards[i], rhs.cards[i]

			if leftCard > rightCard {
				return 1
			}

			if leftCard < rightCard {
				return -1
			}
		}

		return 0
	})

	return plays
}

func parsePlay(line string) Play {
	tokens := strings.Split(line, " ")
	hand, bidStr := tokens[0], tokens[1]

	bid, err := strconv.Atoi(bidStr)
	utils.Check(err)

	cards := getCards(hand)
	cardCount := countCards(hand)

	return Play{
		hand:      hand,
		cards:     cards,
		bid:       bid,
		handType:  getHandType(cardCount),
		cardCount: cardCount,
	}
}

func getCards(hand string) []int {
	cards := make([]int, len(hand))
	runes := []rune(hand)
	for i := 0; i < len(runes); i++ {
		cards[i] = getCard(runes[i])
	}

	return cards
}

func countCards(hand string) []Card {
	cardMap := make(map[int]int)
	for _, cardRune := range hand {
		cardMap[getCard(cardRune)] += 1
	}

	index := 0
	cards := make([]Card, len(cardMap))
	for card, count := range cardMap {
		cards[index] = Card{card: card, count: count}
		index++
	}

	return cards
}

func getCard(cardRune rune) int {
	switch cardRune {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(cardRune - '0')
	}
}

func getHandType(cards []Card) HandType {
	if isFiveOfAKind(cards) {
		return FiveOfAKind
	} else if isFoak(cards) {
		return FourOfAKind
	} else if isFullHouse(cards) {
		return FullHouse
	} else if isToak(cards) {
		return ThreeOfAKind
	} else if isTwoPair(cards) {
		return TwoPair
	} else if isOnePair(cards) {
		return OnePair
	} else {
		return HighCard
	}
}

func isFiveOfAKind(cards []Card) bool {
	return len(cards) == 1
}

func isFoak(cards []Card) bool {
	if len(cards) != 2 {
		return false
	}

	return cards[0].count == 4 || cards[0].count == 1
}

func isFullHouse(cards []Card) bool {
	if len(cards) != 2 {
		return false
	}

	if cards[0].count == 3 {
		return cards[1].count == 2
	}

	if cards[0].count == 2 {
		return cards[1].count == 3
	}

	return false
}

func isToak(cards []Card) bool {
	if len(cards) != 3 {
		return false
	}

	for _, card := range cards {
		if card.count == 3 {
			return true
		}
	}

	return false
}

func isTwoPair(cards []Card) bool {
	if len(cards) != 3 {
		return false
	}

	numPairs := 0
	for _, card := range cards {
		if card.count == 2 {
			numPairs++
		}
	}

	return numPairs == 2
}

func isOnePair(cards []Card) bool {
	return len(cards) == 4
}
