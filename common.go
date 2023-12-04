package aoc2023

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInput() string {
	dat, err := os.ReadFile("./input")
	check(err)

	return string(dat)

}

func WriteOutput(output string, n int) {
	os.WriteFile(fmt.Sprintf("./output-%d", n), []byte(output+"\n"), 0644)
}

func GetLines(input string) []string {
	lines := strings.Split(input, "\n")

	return slices.DeleteFunc(lines, func(line string) bool {
		return line == ""
	})
}

func ParseInt(input string) int {
	res, err := strconv.Atoi(input)
	check(err)

	return res
}

func SliceMap[Input, Output any](input []Input, f func(Input) Output) []Output {
	res := make([]Output, 0, len(input))

	for _, item := range input {
		res = append(res, f(item))
	}

	return res
}

func SliceReduce[Input, Output any](input []Input, f func(Output, Input) Output, seed Output) Output {
	res := seed

	for _, el := range input {
		res = f(res, el)
	}

	return res
}

func Sum[Input int](a int, b int) int {
	return a + b
}
