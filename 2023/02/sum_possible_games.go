package main

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

// only 12 red cubes, 13 green cubes, and 14 blue cubes
var LOADED = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func SumPossibleGames(inputFile []byte) int {
	bytesReader := bytes.NewReader(inputFile)
	bufScanner := bufio.NewScanner(bytesReader)

	inputSum := 0
	for bufScanner.Scan() {
		scannedText := bufScanner.Text()
		if len(scannedText) == 0 {
			continue
		}

		gamePossible := true

		gameIDRounds := strings.Split(scannedText, ":")
		gameID, _ := strconv.Atoi(strings.Split(gameIDRounds[0], " ")[1])

		rounds := strings.Split(gameIDRounds[1], ";")

		for _, round := range rounds {
			colors := strings.Split(round, ",")
			for _, c := range colors {
				c := strings.TrimSpace(c)
				qtyColor := strings.Split(c, " ")

				qty, _ := strconv.Atoi(qtyColor[0])
				color := qtyColor[1]

				if qty > LOADED[color] {
					gamePossible = false
				}
			}
		}

		if gamePossible {
			inputSum += gameID
		}
	}

	return inputSum
}
