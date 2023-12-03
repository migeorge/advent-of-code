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

	partNumberSum, gearPowers := ProcessInput(fileBytes)
	fmt.Println("Part Number Sum:", partNumberSum)
	fmt.Println("Gear Powers:", gearPowers)
}
