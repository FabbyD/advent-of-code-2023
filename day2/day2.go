package main

import (
	"advent-of-code-2023/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const red = 0
const green = 1
const blue = 2

var colors map[string]int = map[string]int{
	"red":   red,
	"green": green,
	"blue":  blue,
}

var cubesConfig [3]int = [3]int{12, 13, 14}

func main() {
	part2()
}

func part1() {
	file, err := os.Open("day2-input.txt")
	utils.Check(err)
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		gameId, draws := parseGame(line)

		isPossibleGame := true
		for _, draw := range draws {
			if !isPossibleDraw(draw) {
				isPossibleGame = false
				break
			}
		}

		if isPossibleGame {
			sum += gameId
		}
	}

	fmt.Printf("Solution : %d\n", sum)
}

func part2() {
	file, err := os.Open("day2-input.txt")
	utils.Check(err)
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		_, draws := parseGame(line)
		maxCubes := findMaxCubes(draws)
		sum += computePower(maxCubes)
	}

	fmt.Printf("Solution : %d\n", sum)
}

func parseDraw(drawString string) []int {
	draw := make([]int, 3)
	tokens := strings.Split(drawString, ",")
	for _, colorString := range tokens {
		colorString = strings.TrimSpace(colorString)
		colorTokens := strings.Split(colorString, " ")
		count, err := strconv.Atoi(colorTokens[0])
		utils.Check(err)
		color := colorTokens[1]
		draw[colors[color]] = count
	}

	return draw
}

func parseGame(line string) (gameId int, draws [][]int) {
	tokens := strings.Split(line, ":")
	gameString := tokens[0]
	drawsString := tokens[1]

	tokens = strings.Split(gameString, " ")
	gameId, err := strconv.Atoi(tokens[1])
	utils.Check(err)

	tokens = strings.Split(drawsString, ";")

	draws = make([][]int, 0)
	for _, drawString := range tokens {
		draw := parseDraw(drawString)
		draws = append(draws, draw)
	}

	return gameId, draws
}

func isPossibleDraw(draw []int) bool {
	for color, count := range draw {
		if cubesConfig[color] < count {
			return false
		}
	}

	return true
}

func findMaxCubes(draws [][]int) []int {
	maxCubes := make([]int, 3)

	for _, draw := range draws {
		for color, count := range draw {
			if maxCubes[color] < count {
				maxCubes[color] = count
			}
		}
	}

	return maxCubes
}

func computePower(cubeset []int) int {
	power := 1
	for _, count := range cubeset {
		power *= count
	}
	return power
}
