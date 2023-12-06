package main

import (
	"advent-of-code-2023/utils"
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	start  uint64
	length uint64
}

type MapRange struct {
	sourceStart uint64
	destStart   uint64
	length      uint64
}

type Mapping struct {
	entries []MapRange
}

func (mapping *Mapping) Map(sourceNumber uint64) uint64 {
	for _, entry := range mapping.entries {
		lastSource := entry.sourceStart + entry.length - 1
		if entry.sourceStart <= sourceNumber && sourceNumber <= lastSource {
			delta := sourceNumber - entry.sourceStart
			return entry.destStart + delta
		}
	}

	return sourceNumber
}

func (mapping *Mapping) MapMany(sourceNumbers []uint64) []uint64 {
	destNumbers := make([]uint64, len(sourceNumbers))
	for i, sourceNumber := range sourceNumbers {
		destNumbers[i] = mapping.Map(sourceNumber)
	}

	return destNumbers
}

func concat(a []Range, b []Range) []Range {
	newArray := make([]Range, len(a)+len(b))

	if len(a) == 1 {
		newArray[0] = a[0]
	} else if len(a) > 0 {
		copy(newArray[:len(a)], a)
	}

	if len(b) == 1 {
		newArray[len(a)] = b[0]
	} else if len(b) > 0 {
		copy(newArray[len(a):], b)
	}

	return newArray
}

func (mapping *Mapping) MapRange(r Range) []Range {
	if r.length == 0 {
		return make([]Range, 0)
	}

	last := r.start + r.length - 1

	mappedRanges := make([]Range, 0)

	for _, entry := range mapping.entries {
		lastSource := entry.sourceStart + entry.length - 1

		if entry.sourceStart <= r.start && last <= lastSource {
			// fully within entry
			delta := r.start - entry.sourceStart
			mappedRange := Range{
				start:  entry.destStart + delta,
				length: r.length,
			}
			mappedRanges = append(mappedRanges, mappedRange)
			return mappedRanges
		}

		if last >= entry.sourceStart && last <= lastSource {
			// overlaps the beginning of entry
			leftRange := Range{
				start:  r.start,
				length: entry.sourceStart - r.start,
			}
			newMappedRanges := mapping.MapRange(leftRange)
			mappedRanges = concat(mappedRanges, newMappedRanges)

			rightRange := Range{
				start:  r.start + leftRange.length,
				length: r.length - leftRange.length,
			}
			newMappedRanges = mapping.MapRange(rightRange)
			mappedRanges = concat(mappedRanges, newMappedRanges)

			return mappedRanges
		}

		if entry.sourceStart <= r.start && r.start < lastSource {
			// overlaps the end of entry
			leftRange := Range{
				start:  r.start,
				length: lastSource - r.start,
			}
			newMappedRanges := mapping.MapRange(leftRange)
			mappedRanges = concat(mappedRanges, newMappedRanges)

			rightRange := Range{
				start:  r.start + leftRange.length,
				length: r.length - leftRange.length,
			}
			newMappedRanges = mapping.MapRange(rightRange)
			mappedRanges = concat(mappedRanges, newMappedRanges)

			return mappedRanges
		}

		if r.start < entry.sourceStart && last > lastSource {
			// overlaps both start and end of entry
			leftRange := Range{
				start:  r.start,
				length: entry.sourceStart - r.start,
			}
			newMappedRanges := mapping.MapRange(leftRange)
			mappedRanges = concat(mappedRanges, newMappedRanges)

			// can just take this whole entry's range and map right away
			centerRange := Range{
				start:  entry.destStart,
				length: entry.length,
			}
			mappedRanges = append(mappedRanges, centerRange)

			rightRange := Range{
				start:  r.start + leftRange.length + centerRange.length,
				length: r.length - leftRange.length - centerRange.length,
			}
			newMappedRanges = mapping.MapRange(rightRange)
			mappedRanges = concat(mappedRanges, newMappedRanges)

			return mappedRanges
		}
	}

	mappedRanges = append(mappedRanges, r)
	return mappedRanges
}

func (mapping *Mapping) MapRanges(ranges []Range) []Range {
	newRanges := make([]Range, 0)
	for _, r := range ranges {
		mappedRanges := mapping.MapRange(r)
		newRanges = concat(newRanges, mappedRanges)
	}

	return newRanges
}

func main() {
	part2()
}

func part1() {
	file := utils.OpenInput()
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	currentNumbers := parseSeeds(scanner.Text())
	fmt.Printf("Seeds : %v\n", currentNumbers)

	scanner.Scan() // skip new line

	for scanner.Scan() { // reads header
		mapping := readMappingBlock(scanner)
		currentNumbers = mapping.MapMany(currentNumbers)
		fmt.Printf("Numbers : %v\n", currentNumbers)
	}

	min := slices.Min(currentNumbers)
	fmt.Printf("Soluton : %v\n", min)
}

func part2() {
	file := utils.OpenInput()
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	currentRanges := readSeedRanges(scanner.Text())
	fmt.Printf("       seeds : %v\n", currentRanges)

	scanner.Scan() // skip new line

	for scanner.Scan() { // reads header
		category := parseDestCategory(scanner.Text())
		mapping := readMappingBlock(scanner)
		currentRanges = mapping.MapRanges(currentRanges)
		fmt.Printf("%12s : %v\n", category, currentRanges)
	}

	comparisonFunc := func(lhs Range, rhs Range) int {
		if lhs.start == rhs.start {
			return 0
		}

		if lhs.start < rhs.start {
			return -1
		}

		return 1
	}

	min := slices.MinFunc(currentRanges, comparisonFunc)
	fmt.Printf("Solution : %v\n", min)
}

func parseDestCategory(s string) string {
	tokens := strings.Split(s, " ")
	tokens = strings.Split(tokens[0], "-")
	return tokens[2]
}

func readSeedRanges(line string) []Range {
	tokens := strings.Split(line, " ")

	numRanges := (len(tokens) - 1) / 2
	ranges := make([]Range, numRanges)
	for i := 1; i < len(tokens); i += 2 {
		start := parseNumber(tokens[i])
		r := Range{
			start:  start,
			length: parseNumber(tokens[i+1]),
		}

		ranges[(i-1)/2] = r
	}

	return ranges
}

func parseSeeds(line string) []uint64 {
	tokens := strings.Split(line, " ")

	seeds := make([]uint64, len(tokens)-1)
	for i := 1; i < len(tokens); i++ {
		seed, err := strconv.ParseUint(tokens[i], 10, 64)
		utils.Check(err)

		seeds[i-1] = seed
	}

	return seeds
}

func readMappingBlock(scanner *bufio.Scanner) Mapping {
	entries := make([]MapRange, 0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		tokens := strings.Split(line, " ")
		r := MapRange{
			destStart:   parseNumber(tokens[0]),
			sourceStart: parseNumber(tokens[1]),
			length:      parseNumber(tokens[2]),
		}

		entries = append(entries, r)
	}

	return Mapping{entries: entries}
}

func parseNumber(s string) uint64 {
	number, err := strconv.ParseUint(s, 10, 64)
	utils.Check(err)

	return number
}
