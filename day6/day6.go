package main

import (
	"advent-of-code-2023/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	time     float64
	distance float64
}

func main() {
	part2()
}

func part1() {
	file := utils.OpenInput()

	races := readRaces(file)

	fmt.Println("Races:", races)

	solution := float64(1)
	for _, race := range races {
		ways := numWays(race)
		solution *= ways
	}

	fmt.Println("Solution:", solution)
}

func part2() {
	file := utils.OpenInput()

	race := readRace(file)

	fmt.Printf("Solution: %f\n", numWays(race))
}

func readRaces(file *os.File) []Race {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timeLine := scanner.Text()

	scanner.Scan()
	distanceLine := scanner.Text()

	r, _ := regexp.Compile("([0-9]+)")

	times := r.FindAllString(timeLine, -1)
	distances := r.FindAllString(distanceLine, -1)

	races := make([]Race, len(times))

	for i := 0; i < len(races); i++ {
		time, err := strconv.Atoi(times[i])
		utils.Check(err)

		distance, err := strconv.Atoi(distances[i])
		utils.Check(err)

		races[i] = Race{
			time:     float64(time),
			distance: float64(distance),
		}
	}

	return races
}

func readRace(file *os.File) Race {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timeLine := scanner.Text()

	scanner.Scan()
	distanceLine := scanner.Text()

	r, _ := regexp.Compile("([0-9])")

	time := parseDigits(r.FindAllString(timeLine, -1))
	distance := parseDigits(r.FindAllString(distanceLine, -1))

	return Race{
		time:     time,
		distance: distance,
	}
}

func parseDigits(digits []string) float64 {
	num := float64(0)
	for _, digitStr := range digits {
		digit, err := strconv.ParseFloat(digitStr, 64)
		utils.Check(err)

		num = num*10 + digit
	}

	return num
}

func numWays(race Race) float64 {
	powTime := math.Pow(race.time, 2)
	sqrt := math.Sqrt(powTime - 4*race.distance)
	minTimeHeld := math.Ceil((race.time - sqrt) / 2)
	maxTimeHeld := math.Floor((race.time + sqrt) / 2)

	if calculateDistance(minTimeHeld, race.time) == race.distance {
		minTimeHeld += 1
	}

	if calculateDistance(maxTimeHeld, race.time) == race.distance {
		maxTimeHeld -= 1
	}

	fmt.Println("Race:", race, "Min:", minTimeHeld, "Max:", maxTimeHeld)

	return maxTimeHeld - minTimeHeld + 1
}

func calculateDistance(timeHeld float64, raceTime float64) float64 {
	return raceTime*timeHeld - math.Pow(timeHeld, 2)
}
