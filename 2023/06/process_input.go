package main

import (
	"strconv"
	"strings"
)

func ProcessInput(inputContents string) (int, int) {
	lines := strings.Split(inputContents, "\n")

	return part1(lines), part2(lines)
}

func part1(lines []string) int {
	TIMES := []int{}
	DISTANCES := []int{}

	product := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		splitLine := strings.Split(line, ":")
		label := splitLine[0]
		values := splitLine[1]

		valArr := strings.Split(values, " ")

		var target *[]int
		if strings.ToLower(label) == "time" {
			target = &TIMES
		}
		if strings.ToLower(label) == "distance" {
			target = &DISTANCES
		}

		for _, val := range valArr {
			if len(val) == 0 {
				continue
			}

			valInt, _ := strconv.Atoi(val)
			*target = append(*target, valInt)
		}
	}

	for i, raceTime := range TIMES {
		distanceRecord := DISTANCES[i]
		possibleWins := NumPossibleWins(raceTime, distanceRecord)

		if product == 0 {
			product = possibleWins
			continue
		}

		product *= possibleWins
	}

	return product
}

func part2(lines []string) int {
	var TIME int
	var DISTANCE int

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		splitLine := strings.Split(line, ":")
		label := splitLine[0]
		values := splitLine[1]

		valStr := strings.ReplaceAll(values, " ", "")

		var target *int
		if strings.ToLower(label) == "time" {
			target = &TIME
		}
		if strings.ToLower(label) == "distance" {
			target = &DISTANCE
		}

		valInt, _ := strconv.Atoi(valStr)
		*target = valInt
	}

	return NumPossibleWins(TIME, DISTANCE)
}

func NumPossibleWins(raceTime, distanceRecord int) int {
	possibleWins := 0

	for i := 0; i < raceTime; i++ {
		timeLeft := raceTime - i
		speed := i

		distance := timeLeft * speed

		if distance > distanceRecord {
			possibleWins++
		}
	}

	return possibleWins
}
