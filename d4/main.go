package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unumtresocto/aoc2023"
)

type Card struct {
	numbers []int
	winners []int
}

func trimDelimiters(line string) string {
	return strings.TrimFunc(line, func(r rune) bool {
		return r == ':' || r == '|' || r == ' '
	})
}

func parseNumbers(line string) []int {
	var numbers []int

	strings := regexp.MustCompile("\\s+").Split(line, -1)

	for i := range strings {
		numbers = append(numbers, aoc2023.ParseInt(strings[i]))
	}

	return numbers
}

func parseLine(line string) Card {
	winnersRe := regexp.MustCompile(":.+\\|")
	numbersRe := regexp.MustCompile("\\|.+$")

	winnersString := trimDelimiters(winnersRe.FindString(line))
	numbersString := trimDelimiters(numbersRe.FindString(line))

	sortedWinners := parseNumbers(winnersString)
	slices.Sort(sortedWinners)

	return Card{
		numbers: parseNumbers(numbersString),
		winners: sortedWinners,
	}
}

const (
	MODE_4_1 = 1
	MODE_4_2 = 2
)

func getScore(card Card, mode int) int {
	res := 0

	for _, num := range card.numbers {
		if slices.Index(card.winners, num) != -1 {
			if res == 0 || mode == MODE_4_2 {
				res += 1
			} else {
				res *= 2
			}
		}
	}

	return res
}

func getMultipliedCards(cards []Card) int {
	copies := make([]int, len(cards), len(cards))

	for i := range copies {
		copies[i] = 1
	}

	for i, card := range cards {
		score := getScore(card, 2)

		for j := 0; j < score; j++ {
			copies[i+1+j] += copies[i]
		}
	}

	return aoc2023.SliceReduce(copies, aoc2023.Sum, 0)
}

func getScoreModeOne(card Card) int {
	return getScore(card, MODE_4_1)
}

func main() {
	input := aoc2023.ReadInput()
	lines := aoc2023.GetLines(input)
	cards := aoc2023.SliceMap(lines, parseLine)
	scores := aoc2023.SliceMap(cards, getScoreModeOne)

	result := aoc2023.SliceReduce(scores, aoc2023.Sum, 0)
	result2 := getMultipliedCards(cards)

	aoc2023.WriteOutput(fmt.Sprint(result), 1)
	aoc2023.WriteOutput(fmt.Sprint(result2), 2)
}
