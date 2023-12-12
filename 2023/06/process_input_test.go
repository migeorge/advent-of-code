package main

import (
	"testing"
)

const SAMPLE_INPUT = `
Time:      7  15   30
Distance:  9  40  200
`

func TestProcessInput(t *testing.T) {
	part1, part2 := ProcessInput(SAMPLE_INPUT)
	expectedWinProductPart1 := 288
	expectedWinProductPart2 := 71503

	if part1 != expectedWinProductPart1 || part2 != expectedWinProductPart2 {
		t.Fatalf(`ProcessInput(SAMPLE_INPUT) = %d, %d want %d, %d`, part1, part2, expectedWinProductPart1, expectedWinProductPart2)
	}
}

func TestNumPossibleWins(t *testing.T) {
	type numPossibleWinsTestCase struct {
		raceTime        int
		distanceRecord  int
		expectedNumWins int
	}

	for _, testCase := range []numPossibleWinsTestCase{
		{
			raceTime:        7,
			distanceRecord:  9,
			expectedNumWins: 4,
		},
		{
			raceTime:        71530,
			distanceRecord:  940200,
			expectedNumWins: 71503,
		},
		{
			raceTime:        53916768,
			distanceRecord:  250133010811025,
			expectedNumWins: 43663323,
		},
	} {
		numWins := NumPossibleWins(testCase.raceTime, testCase.distanceRecord)
		if numWins != testCase.expectedNumWins {
			t.Errorf(`NumPossibleWins(%d, %d) = %d want %d`, testCase.raceTime, testCase.distanceRecord, numWins, testCase.expectedNumWins)
		}
	}
}
