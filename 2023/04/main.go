package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	inputPath := flag.String("input", "", "input file path")
	flag.Parse()

	if len(*inputPath) == 0 {
		fmt.Println("input flag must be provided")
		os.Exit(1)
	}

	fileBytes, err := os.ReadFile(*inputPath)
	if err != nil {
		fmt.Println("unable to read input file:", err.Error())
		os.Exit(1)
	}

	pointSum, totalNumCards := ProcessInput(fileBytes)
	fmt.Println("Point Sum:", pointSum)
	fmt.Println("Total Num Cards:", totalNumCards)
}
