package main

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var INPUT = []string{}
var SYMBOL_NUMS = map[string][]int{}

var NUM_REG = regexp.MustCompile(`[0-9]+`)

const NOT_SYMBOL = ".0123456789"

func ProcessInput(inputFile []byte) (int, int) {
	bytesReader := bytes.NewReader(inputFile)
	bufScanner := bufio.NewScanner(bytesReader)

	inputSum := 0

	INPUT = append(INPUT, "")
	for bufScanner.Scan() {
		scannedText := bufScanner.Text()
		if len(scannedText) == 0 {
			continue
		}

		INPUT = append(INPUT, "."+scannedText+".")
	}
	INPUT = append(INPUT, "")

	lineLength := len(INPUT[1])
	paddingStr := ""
	for i := 0; i < lineLength; i++ {
		paddingStr += "."
	}

	INPUT[0] = paddingStr
	INPUT[len(INPUT)-1] = paddingStr

	for lineIndex, line := range INPUT {
		foundNumStrs := NUM_REG.FindAllString(line, -1)
		foundNumIndexes := NUM_REG.FindAllStringIndex(line, -1)

		for i, foundNumStr := range foundNumStrs {
			foundNum, _ := strconv.Atoi(foundNumStr)
			numIndexes := foundNumIndexes[i]

			startIndex := numIndexes[0] - 1
			endIndex := numIndexes[1] + 1

			for lineToSearch := lineIndex - 1; lineToSearch < lineIndex+2; lineToSearch++ {
				for j := startIndex; j < endIndex; j++ {
					char := string(INPUT[lineToSearch][j])
					if !strings.Contains(NOT_SYMBOL, char) {
						// hit symbol
						symName := fmt.Sprintf("%s:%d,%d", char, lineToSearch, j)

						nums, exists := SYMBOL_NUMS[symName]
						if exists {
							nums = append(nums, foundNum)
						} else {
							nums = []int{foundNum}
						}
						SYMBOL_NUMS[symName] = nums
					}
				}
			}
		}
	}

	powerSum := 0
	for key, nums := range SYMBOL_NUMS {
		symChar := strings.Split(key, ":")[0]
		var potentialGear bool
		if symChar == "*" {
			potentialGear = true
		}

		var power int
		for _, num := range nums {
			inputSum += num

			if potentialGear && len(nums) == 2 {
				if power == 0 {
					power = num
				} else {
					power *= num
				}
			}
		}
		powerSum += power
	}

	return inputSum, powerSum
}
