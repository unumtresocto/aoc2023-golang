package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var digits []string = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var digitsStr string = strings.Join(digits, "|")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func leftMatch(text string) string {
	index := -1

	for i := 0; i < len(text) && index == -1; i++ {
		substr := text[0 : i+1]
		index = slices.IndexFunc(digits, func(dig string) bool {
			return strings.Index(substr, dig) != -1
		})
	}

	return digits[index]
}

func rightMatch(text string) string {
	index := -1
	textLen := len(text)

	for i := 0; i < textLen && index == -1; i++ {
		substr := text[textLen-1-i : textLen]
		index = slices.IndexFunc(digits, func(dig string) bool {
			return strings.Index(substr, dig) != -1
		})
	}

	return digits[index]
}

func toInt(text string) int {
	res, err := strconv.Atoi(text)

	if err != nil {
		res = slices.Index[[]string](digits, text)
	}

	return res
}

func main() {
	dat, err := os.ReadFile("./input")
	check(err)

	lines := strings.Split(string(dat), "\n")
	result := 0

	for i := 0; i < len(lines)-1; i++ {
		firstDigit := leftMatch(lines[i])
		lastDigit := rightMatch(lines[i])
		num := toInt(firstDigit)*10 + toInt(lastDigit)
		result = result + num
	}

	os.WriteFile("./output", []byte(fmt.Sprint(result)+"\n"), 0644)
}
