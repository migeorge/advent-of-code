package main

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

func PowerMinGames(inputFile []byte) int {
	bytesReader := bytes.NewReader(inputFile)
	bufScanner := bufio.NewScanner(bytesReader)

	inputSum := 0
	for bufScanner.Scan() {
		scannedText := bufScanner.Text()
		if len(scannedText) == 0 {
			continue
		}

		gameIDRounds := strings.Split(scannedText, ":")
		rounds := strings.Split(gameIDRounds[1], ";")

		redMax, greenMax, blueMax := 0, 0, 0

		for _, round := range rounds {
			colors := strings.Split(round, ",")

			for _, c := range colors {
				c := strings.TrimSpace(c)
				qtyColor := strings.Split(c, " ")

				qty, _ := strconv.Atoi(qtyColor[0])
				color := qtyColor[1]

				switch color {
				case "red":
					if qty > redMax {
						redMax = qty
					}
				case "green":
					if qty > greenMax {
						greenMax = qty
					}
				case "blue":
					if qty > blueMax {
						blueMax = qty
					}
				}
			}
		}

		inputSum += redMax * greenMax * blueMax
	}

	return inputSum
}
