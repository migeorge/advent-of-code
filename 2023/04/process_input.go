package main

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

type Card struct {
	ID             int
	Numbers        []int
	WinningNumbers []int
	Points         int
}

var CARDS []*Card
var WON_CARDS []*Card

var CARD_MAP = map[int]*Card{}

func ProcessInput(inputFile []byte) (int, int) {
	bytesReader := bytes.NewReader(inputFile)
	bufScanner := bufio.NewScanner(bytesReader)

	for bufScanner.Scan() {
		scannedText := bufScanner.Text()
		if len(scannedText) == 0 {
			continue
		}

		idValSplit := strings.Split(scannedText, ":")
		idParts := strings.Split(idValSplit[0], " ")
		idStr := idParts[len(idParts)-1]
		id, _ := strconv.Atoi(idStr)

		card := &Card{
			ID:             id,
			Numbers:        []int{},
			WinningNumbers: []int{},
		}

		val := strings.Split(idValSplit[1], "|")

		winningNumberStrs := strings.Split(strings.TrimSpace(val[0]), " ")
		numberStrs := strings.Split(strings.TrimSpace(val[1]), " ")

		for _, num := range winningNumberStrs {
			numI, err := strconv.Atoi(num)
			if err != nil {
				continue
			}
			card.WinningNumbers = append(card.WinningNumbers, numI)
		}
		for _, num := range numberStrs {
			numI, _ := strconv.Atoi(num)
			card.Numbers = append(card.Numbers, numI)
		}

		CARDS = append(CARDS, card)
		CARD_MAP[card.ID] = card
	}

	pointSum := 0
	for _, card := range CARDS {
		calculatePointsAndWonCards(card, true)
		pointSum += card.Points
	}

	totalNumCards := len(CARDS) + len(WON_CARDS)
	return pointSum, totalNumCards
}

func calculatePointsAndWonCards(card *Card, isRootCall bool) {
	cardPoints := 0
	cardNumWins := 0
	for _, cardNum := range card.Numbers {
		for _, winningCardNum := range card.WinningNumbers {
			if cardNum == winningCardNum {
				cardNumWins++

				// To maintain the point count for part 1, don't calculate points for won cards
				if isRootCall {
					if cardPoints == 0 {
						cardPoints = 1
						break
					}

					cardPoints *= 2
				}

				break
			}
		}
	}
	card.Points = cardPoints

	for i := (card.ID + 1); i <= (card.ID + cardNumWins); i++ {
		wonCard := CARD_MAP[i]
		WON_CARDS = append(WON_CARDS, wonCard)
		calculatePointsAndWonCards(wonCard, false)
	}
}
