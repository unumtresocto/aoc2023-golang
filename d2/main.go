package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Round struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id     int
	rounds []Round
}

type Limits struct {
	red   int
	green int
	blue  int
}

func parseGame(gameStr string) Game {
	idRe := regexp.MustCompile("\\d+:")
	roundsRe := regexp.MustCompile(":.+")
	id, _ := strconv.Atoi(strings.Trim(idRe.FindString(gameStr), ":"))
	roundLines := strings.Split(strings.Trim(roundsRe.FindString(gameStr), ":"), ";")

	rounds := []Round{}
	for i := 0; i < len(roundLines); i++ {
		rounds = append(rounds, parseRound(roundLines[i]))
	}

	return Game{id, rounds}
}

func parseRound(roundStr string) Round {
	redRe := regexp.MustCompile("\\d+ red")
	greenRe := regexp.MustCompile("\\d+ green")
	blueRe := regexp.MustCompile("\\d+ blue")

	red, _ := strconv.Atoi(strings.Trim(redRe.FindString(roundStr), " red"))
	green, _ := strconv.Atoi(strings.Trim(greenRe.FindString(roundStr), " green"))
	blue, _ := strconv.Atoi(strings.Trim(blueRe.FindString(roundStr), " blue"))

	return Round{red, green, blue}
}

func isValidGame(game Game, limits Limits) bool {
	isValid := true

	for i := 0; isValid && i < len(game.rounds); i++ {
		round := game.rounds[i]

		if round.red > limits.red || round.green > limits.green || round.blue > limits.blue {
			isValid = false
		}
	}

	return isValid
}

func getGamePower(game Game) int {
	minimums := Round{1, 1, 1}

	for i := range game.rounds {
		round := game.rounds[i]

		if minimums.red < round.red {
			minimums.red = round.red
		}
		if minimums.green < round.green {
			minimums.green = round.green
		}
		if minimums.blue < round.blue {
			minimums.blue = round.blue
		}
	}

	return minimums.red * minimums.green * minimums.blue
}

func main() {
	dat, err := os.ReadFile("./input")
	check(err)

	rawLines := strings.Split(string(dat), "\n")
	gameLines := rawLines[:len(rawLines)-1]

	limits := Limits{12, 13, 14}
	result := 0
	result2 := 0

	for i := range gameLines {
		game := parseGame(gameLines[i])
		delta := getGamePower(game)
		result2 += delta

		if isValidGame(game, limits) {
			result += game.id
		}
	}

	os.WriteFile("./output", []byte(fmt.Sprint(result)+"\n"), 0644)
	os.WriteFile("./output2", []byte(fmt.Sprint(result2)+"\n"), 0644)
}
