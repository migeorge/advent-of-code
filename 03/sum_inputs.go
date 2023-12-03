package main

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

const (
	NON_SYMBOL = "."
	DIGITS     = "1234567890"
)

var LINES = []string{}
var SYMBOL_GRID = [][]int{}
var NUMBER_GRID = [][]int{}
var HIT_GRID = [][]int{}

func SumInputs(inputFile []byte) int {
	bytesReader := bytes.NewReader(inputFile)
	bufScanner := bufio.NewScanner(bytesReader)

	inputSum := 0
	for bufScanner.Scan() {
		scannedText := bufScanner.Text()
		if len(scannedText) == 0 {
			continue
		}

		LINES = append(LINES, scannedText)

		symbolIndexes := []int{}
		for i, c := range scannedText {
			char := string(c)
			if !(strings.Contains(NON_SYMBOL, char) || strings.Contains(DIGITS, char)) {
				symbolIndexes = append(symbolIndexes, i)
			}
		}
		SYMBOL_GRID = append(SYMBOL_GRID, symbolIndexes)

		numberIndexes := []int{}
		for i, c := range scannedText {
			numOffset := c - 48
			if numOffset > -1 && numOffset < 10 {
				numberIndexes = append(numberIndexes, i)
			}
		}
		NUMBER_GRID = append(NUMBER_GRID, numberIndexes)
	}

	for i := range LINES {
		lineHits := []int{}
		for _, numIndex := range NUMBER_GRID[i] {
			var hit bool
			// check current line
			if numIndex != 0 {
				for _, symbolIndex := range SYMBOL_GRID[i] {
					if numIndex-1 == symbolIndex || numIndex+1 == symbolIndex {
						hit = true
					}
				}
			}
			// check line above
			if i > 0 {
				for _, symbolIndex := range SYMBOL_GRID[i-1] {
					if numIndex-1 == symbolIndex || numIndex == symbolIndex || numIndex+1 == symbolIndex {
						hit = true
					}
				}
			}
			// check line below
			if i < len(LINES)-1 {
				for _, symbolIndex := range SYMBOL_GRID[i+1] {
					if numIndex-1 == symbolIndex || numIndex == symbolIndex || numIndex+1 == symbolIndex {
						hit = true
					}
				}
			}

			if hit {
				lineHits = append(lineHits, numIndex)
			}
		}
		HIT_GRID = append(HIT_GRID, lineHits)
	}

	numReg := regexp.MustCompile(`[0-9]+`)

	for i, line := range LINES {
		lineHits := HIT_GRID[i]
		if len(lineHits) > 0 {
			strMatch := numReg.FindAllString(line, -1)
			strMatchIndex := numReg.FindAllStringIndex(line, -1)

			for j, match := range strMatch {
				matchHit := false
				for _, hitIndex := range lineHits {
					if hitIndex == strMatchIndex[j][0] || hitIndex == strMatchIndex[j][1] || (hitIndex > strMatchIndex[j][0] && hitIndex < strMatchIndex[j][1]) {
						matchHit = true
					}
				}

				if matchHit {
					matchNum, _ := strconv.Atoi(match)
					inputSum += matchNum
				}
			}
		}
	}

	return inputSum
}
