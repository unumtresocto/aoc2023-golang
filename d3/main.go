package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Range struct {
	from int
	to   int
}

type Pair struct {
	x int
	y int
}

func processLine(lines []string, i int, gears map[Pair][]int) int {
	numRe := regexp.MustCompile("\\d+")
	ranges := numRe.FindAllStringIndex(lines[i], -1)
	result := 0

	for j := range ranges {
		rng := Range{ranges[j][0], ranges[j][1]}

		if isPartNumber(lines, rng, i, gears) {
			delta := rangeToInt(lines, i, rng)
			result += delta
		}
	}

	return result
}

func rangeToInt(lines []string, lineNumber int, rng Range) int {
	result, _ := strconv.Atoi(lines[lineNumber][rng.from:rng.to])

	return result
}

func isPartNumber(lines []string, rng Range, lineNumber int, gears map[Pair][]int) bool {
	perimeter := getPerimeter(rng, lineNumber)
	validPerimeter := slices.DeleteFunc(perimeter, func(pair Pair) bool {
		return !(pair.x >= 0 && pair.x < len(lines[0]) && pair.y >= 0 && pair.y < len(lines))
	})

	gearLocs := slices.DeleteFunc(validPerimeter, func(pair Pair) bool {
		return string(lines[pair.y][pair.x]) != "*"
	})

	for i := range gearLocs {
		gears[gearLocs[i]] = append(gears[gearLocs[i]], rangeToInt(lines, lineNumber, rng))
	}

	return len(slices.DeleteFunc(validPerimeter, func(pair Pair) bool {
		return string(lines[pair.y][pair.x]) == "."
	})) != 0
}

func getPerimeter(rng Range, index int) []Pair {
	var res []Pair

	res = append(res, Pair{rng.from - 1, index - 1}, Pair{rng.to, index - 1}, Pair{rng.from - 1, index + 1}, Pair{rng.to, index + 1})
	res = append(res, Pair{rng.from - 1, index}, Pair{rng.to, index})

	for i := rng.from; i < rng.to; i++ {
		res = append(res, Pair{i, index - 1}, Pair{i, index + 1})
	}

	return res
}

func main() {
	dat, err := os.ReadFile("./input")
	check(err)

	rawLines := strings.Split(string(dat), "\n")
	lines := rawLines[:len(rawLines)-1]
	gears := make(map[Pair][]int, 0)

	result := 0
	result2 := 0

	for i := range lines {
		result += processLine(lines, i, gears)
	}

	for _, ratios := range gears {
		if len(ratios) == 2 {
			result2 += ratios[0] * ratios[1]
		}
	}

	os.WriteFile("./output", []byte(fmt.Sprint(result)+"\n"), 0644)
	os.WriteFile("./output2", []byte(fmt.Sprint(result2)+"\n"), 0644)
}
