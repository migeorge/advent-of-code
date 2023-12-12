package main

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

var TEXT_NUM = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func SumInputs(inputFile []byte) int {
	bytesReader := bytes.NewReader(inputFile)
	bufScanner := bufio.NewScanner(bytesReader)

	numRegex := regexp.MustCompile(`[1-9]`)

	inputSum := 0
	for bufScanner.Scan() {
		scannedText := bufScanner.Text()
		if len(scannedText) == 0 {
			continue
		}

		var firstTextNum string
		var lastTextNum string
		firstTextNumIndex := -1
		lastTextNumIndex := -1
		for text := range TEXT_NUM {
			firstIndex := strings.Index(scannedText, text)
			lastIndex := strings.LastIndex(scannedText, text)

			if firstIndex != -1 {
				if firstTextNumIndex == -1 || (firstIndex < firstTextNumIndex) {
					firstTextNumIndex = firstIndex
					firstTextNum = text
				}
			}

			if lastIndex != -1 {
				if lastTextNumIndex == -1 || (lastIndex > lastTextNumIndex) {
					lastTextNumIndex = lastIndex
					lastTextNum = text
				}
			}
		}

		replaceTextNum := scannedText
		if len(firstTextNum) > 0 {
			replaceTextNum = strings.Replace(replaceTextNum, firstTextNum, TEXT_NUM[firstTextNum], 1)
		}
		if len(lastTextNum) > 0 {
			replaceTextNum = strings.ReplaceAll(replaceTextNum, lastTextNum, TEXT_NUM[lastTextNum])
		}

		numString := numRegex.FindAllString(replaceTextNum, -1)

		firstNum := numString[0]
		lastNum := numString[len(numString)-1]

		combinedNum, _ := strconv.Atoi(firstNum + lastNum)

		inputSum += combinedNum
	}

	return inputSum
}
