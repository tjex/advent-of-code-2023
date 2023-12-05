package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Game struct {
	id, red, green, blue int
}

type ColorResultPerGame struct {
	red, green, blue []int
}

var gameData []string
var redCountPerGame []string
var greenCountPerGame []string
var blueCountPerGame []string

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

	ids, countsMax, _ := getGameData(gameData)
	redMax := countsMax.red
	greenMax := countsMax.green
	blueMax := countsMax.blue

	partOne(ids, redMax, greenMax, blueMax)
	partTwo(ids, redMax, greenMax, blueMax)

}

func partOne(ids, redMax, greenMax, blueMax []int) {
	answer := 0
	for i := range ids {
		game := generateStruct(ids[i], redMax[i], greenMax[i], blueMax[i])
		// check if game was possible
		if game.red <= 12 && game.green <= 13 && game.blue <= 14 {
			answer += game.id
		}
	}
	fmt.Println("part 1:", answer)
}

func partTwo(ids, redMax, greenMax, blueMax []int) {
	answer := 0
	for i := range ids {
		game := generateStruct(ids[i], redMax[i], greenMax[i], blueMax[i])

        // fmt.Println(game.red, game.green, game.blue)
        // multiply lowest number of each color together
        power := game.red * game.green * game.blue
        answer += power
	}
	fmt.Println("part 2:", answer)

}

// generate a struct for each game
func generateStruct(id, redCount, greenCount, blueCount int) Game {
	gameStruct := Game{id, redCount, greenCount, blueCount}
	return gameStruct
}

func getLargestNumber(game []string, r *regexp.Regexp) int {
	max := 0
	for _, num := range game {
		n := r.FindAllString(num, -1)
		nAsInt, err := strconv.Atoi(n[0])
		if err != nil {
			fmt.Println(err)
		}
		if nAsInt > max {
			max = nAsInt
		}
	}
	return max
}

func getSmallestNumber(game []string, r *regexp.Regexp) int {
	min := math.MaxInt32
	for _, num := range game {
		n := r.FindAllString(num, -1)
		nAsInt, err := strconv.Atoi(n[0])
		if err != nil {
			fmt.Println(err)
		}
		if nAsInt < min {
			min = nAsInt
		}
	}
	return min
}

func extractIDs(line string) int {
	reID := regexp.MustCompile(`(\d*:)`)
	digit := regexp.MustCompile(`\d*`)
	id := reID.FindString(line)
	id = digit.FindString(id)
	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
		}
		return idInt
	}
	return 0
}

func getGameData(gameData []string) ([]int, ColorResultPerGame, ColorResultPerGame) {
	reRed := regexp.MustCompile(redCountPat)
	reGreen := regexp.MustCompile(greenCountPat)
	reBlue := regexp.MustCompile(blueCountPat)
	digit := regexp.MustCompile(`\d*`)

	var redCountsMax []int
	var greenCountsMax []int
	var blueCountsMax []int
	var redCountsMin []int
	var greenCountsMin []int
	var blueCountsMin []int
	var gameStructIDs []int

	// generate ids string->int for structs
	// (used for final answer generation)
	for _, line := range gameData {

		id := extractIDs(line)
		// a game id of 0 is a blank line in input game data
		if id != 0 {
			gameStructIDs = append(gameStructIDs, id)
		}

		// find color counts
		redCountPerGame = reRed.FindAllString(line, -1)
		greenCountPerGame = reGreen.FindAllString(line, -1)
		blueCountPerGame = reBlue.FindAllString(line, -1)

		// accumulate max counts across all games per color
		redMax := getLargestNumber(redCountPerGame, digit)
		redCountsMax = append(redCountsMax, redMax)

		greenMax := getLargestNumber(greenCountPerGame, digit)
		greenCountsMax = append(greenCountsMax, greenMax)

		blueMax := getLargestNumber(blueCountPerGame, digit)
		blueCountsMax = append(blueCountsMax, blueMax)

		// accumulate min counts across all games per color
		redMin := getSmallestNumber(redCountPerGame, digit)
		redCountsMin = append(redCountsMin, redMin)

		greenMin := getSmallestNumber(greenCountPerGame, digit)
		greenCountsMin = append(greenCountsMin, greenMin)

		blueMin := getSmallestNumber(blueCountPerGame, digit)
		blueCountsMin = append(blueCountsMin, blueMin)
	}

    // fill and return structs instead of arrays to reduce ambiguity when accessing
    // the correct values. e.g. Struct.green instead of array[1]
	Max := ColorResultPerGame{redCountsMax, greenCountsMax, blueCountsMax}
	Min := ColorResultPerGame{redCountsMin, greenCountsMin, blueCountsMin}

	return gameStructIDs, Max, Min
}
