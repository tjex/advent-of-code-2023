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

	ids, redMax, greenMax, blueMax := getGameData(gameData)

	// generate structs
	answer := 0
	for i := range ids {
		game := generateStruct(ids[i], redMax[i], greenMax[i], blueMax[i])
		// check if game was possible
		if game.red <= 12 && game.green <= 13 && game.blue <= 14 {
			answer += game.id
			fmt.Println(game.id)
		}
	}
	fmt.Println(answer)

}

// generate a struct for each line in input game data
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

func getGameData(gameData []string) ([]int, []int, []int, []int) {
	reRed := regexp.MustCompile(redCountPat)
	reGreen := regexp.MustCompile(greenCountPat)
	reBlue := regexp.MustCompile(blueCountPat)
	reID := regexp.MustCompile(`(\d*:)`)
	digit := regexp.MustCompile(`\d*`)

	var redCounts []int
	var greenCounts []int
	var blueCounts []int
	var gameStructIDs []int

	for _, line := range gameData {
		// n := fmt.Sprintf("%s", "game"+strconv.Itoa(i+1))
		id := reID.FindString(line)
		id = digit.FindString(id)
		if id != "" {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(err)
			}
			gameStructIDs = append(gameStructIDs, idInt)
		}

		// find color counts
		redCountPerGame = reRed.FindAllString(line, -1)
		greenCountPerGame := reGreen.FindAllString(line, -1)
		blueCountPerGame := reBlue.FindAllString(line, -1)

		// accumulate counts across all games
		redMax := getLargestNumber(redCountPerGame, digit)
		redCounts = append(redCounts, redMax)

		greenMax := getLargestNumber(greenCountPerGame, digit)
		greenCounts = append(greenCounts, greenMax)

		blueMax := getLargestNumber(blueCountPerGame, digit)
		blueCounts = append(blueCounts, blueMax)

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
