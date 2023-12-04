package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Game struct {
	id, red, green, blue int
}

var gameData []string

const gmeIdPat = `(\d*:)`
const redCountPat = `(\d*\ red)`
const greenCountPat = `(\d*\ green)`
const blueCountPat = `(\d*\ blue)`

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("file open error:", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		gameData = append(gameData, sc.Text())
	}

    ids, redMax, greenMax, blueMax := getGameData(gameData)

	// generate structs
	answer := 0
	for i := range ids {
		game := generateStruct(ids[i], redMax[i], greenMax[i], blueMax[i])
		// fmt.Println(game)
		// check if game was possible
		if game.red <= 12 && game.blue <= 14 && game.green <= 13 {
			answer += game.id
		}
	}
	fmt.Println(answer)

}

// generate a struct for each line in input game data
func generateStruct(id, redCount, greenCount, blueCount int) Game {
	gameStruct := Game{id, redCount, greenCount, blueCount}
	return gameStruct
}

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func getGameData(gameData []string) ([]int, []int, []int, []int) {
	reRed := regexp.MustCompile(redCountPat)
	reGreen := regexp.MustCompile(greenCountPat)
	reBlue := regexp.MustCompile(blueCountPat)
	digit := regexp.MustCompile(`\d*`)

	var redCounts []int
	var greenCounts []int
	var blueCounts []int
	var gameStructIDs []int

	var redCountPerGame []string

	for i, line := range gameData {
		// n := fmt.Sprintf("%s", "game"+strconv.Itoa(i+1))
		gameStructIDs = append(gameStructIDs, i+1)

		// find color counts
		redCountPerGame = reRed.FindAllString(line, -1)
		greenCountPerGame := reGreen.FindAllString(line, -1)
		blueCountPerGame := reBlue.FindAllString(line, -1)

		// accumulate counts across all games
		// for red
		maxRed := 0
		for _, num := range redCountPerGame {
			n := digit.FindAllString(num, -1)
			nAsInt, err := strconv.Atoi(n[0])
			if err != nil {
				fmt.Println(err)
			}
			if nAsInt > maxRed {
				maxRed = nAsInt
			}
		}
		redCounts = append(redCounts, maxRed)

		// for green
		maxGreen := 0
		for _, count := range greenCountPerGame {
			n := digit.FindAllString(count, -1)
			nAsInt, err := strconv.Atoi(n[0])
			if err != nil {
				fmt.Println(err)
			}
			if nAsInt > maxGreen {
				maxGreen = nAsInt
			}

		}
		greenCounts = append(greenCounts, maxRed)

		// for "blue"
		maxBlue := 0
		for _, count := range blueCountPerGame {
			n := digit.FindAllString(count, -1)
			nAsInt, err := strconv.Atoi(n[0])
			if err != nil {
				fmt.Println(err)
			}
			if nAsInt > maxBlue {
				maxBlue = nAsInt
			}
		}
		blueCounts = append(blueCounts, maxGreen)

	}

	return gameStructIDs, redCounts, greenCounts, blueCounts

	// source per line: id, total red, green and blue
	// convert number strings to ints
	// get sum of each color
	// create struct with id, sum red, sum, green, sum blue

}

// filter and return games based on requirements
// return the structs and their ids
func filterGames(req []string) {
}
