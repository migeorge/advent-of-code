package main

import (
	"strconv"
	"strings"
)

type SrcDestRange struct {
	SrcStart   int
	SrcEnd     int
	DestOffest int
}

type SeedRange struct {
	SeedStart int
	SeedEnd   int
}

type Rules []SrcDestRange

func (r *Rules) Destination(source int) int {
	for _, sdr := range *r {
		if source >= sdr.SrcStart && source <= sdr.SrcEnd {
			return source + sdr.DestOffest
		}
	}

	return source
}

type RuleDB struct {
	Seeds                 []int
	SeedRanges            []SeedRange
	SeedToSoil            Rules
	SoilToFertilizer      Rules
	FertilizerToWater     Rules
	WaterToLight          Rules
	LightToTemperature    Rules
	TemperatureToHumidity Rules
	HumidityToLocation    Rules
}

func ProcessInput(inputContents string) (int, int) {
	lines := strings.Split(inputContents, "\n")
	inputChunks := chunkInput(lines)

	part1MinLocation := part1(inputChunks)
	part2MinLocation := part2(inputChunks)

	return part1MinLocation, part2MinLocation
}

func part1(inputChunks [][]string) int {
	ruleDB := &RuleDB{}
	populateSeedsPart1(inputChunks, ruleDB)

	return findMinLocation(inputChunks, ruleDB)
}

func part2(inputChunks [][]string) int {
	ruleDB := &RuleDB{}
	populateSeedsPart2(inputChunks, ruleDB)

	return findMinLocation(inputChunks, ruleDB)
}

func findMinLocation(inputChunks [][]string, ruleDB *RuleDB) int {
	mapInputs(inputChunks, ruleDB)

	minLocation := 0

	if len(ruleDB.Seeds) != 0 {
		for _, seed := range ruleDB.Seeds {
			soil := ruleDB.SeedToSoil.Destination(seed)
			fertilizer := ruleDB.SoilToFertilizer.Destination(soil)
			water := ruleDB.FertilizerToWater.Destination(fertilizer)
			light := ruleDB.WaterToLight.Destination(water)
			temperature := ruleDB.LightToTemperature.Destination(light)
			humidity := ruleDB.TemperatureToHumidity.Destination(temperature)
			location := ruleDB.HumidityToLocation.Destination(humidity)

			if minLocation == 0 || location < minLocation {
				minLocation = location
			}
		}
	} else {
		for _, seedRange := range ruleDB.SeedRanges {
			for seed := seedRange.SeedStart; seed < seedRange.SeedEnd; seed++ {
				soil := ruleDB.SeedToSoil.Destination(seed)
				fertilizer := ruleDB.SoilToFertilizer.Destination(soil)
				water := ruleDB.FertilizerToWater.Destination(fertilizer)
				light := ruleDB.WaterToLight.Destination(water)
				temperature := ruleDB.LightToTemperature.Destination(light)
				humidity := ruleDB.TemperatureToHumidity.Destination(temperature)
				location := ruleDB.HumidityToLocation.Destination(humidity)

				if minLocation == 0 || location < minLocation {
					minLocation = location
				}
			}
		}
	}

	return minLocation
}

func chunkInput(lines []string) [][]string {
	inputChunks := [][]string{}

	inputChunk := []string{}
	for _, line := range lines {
		if len(line) == 0 {
			inputChunks = append(inputChunks, inputChunk)
			inputChunk = []string{}
			continue
		}

		inputChunk = append(inputChunk, line)
	}

	return inputChunks
}

func populateSeedsPart1(chunks [][]string, ruleDB *RuleDB) {
	var seeds []int

	for _, chunk := range chunks {
		if len(chunk) == 0 {
			continue
		}

		if len(chunk) == 1 {
			chunkStr := chunk[0]

			if seedsStr, found := strings.CutPrefix(chunkStr, "seeds: "); found {
				seedsStrs := strings.Split(seedsStr, " ")
				seeds = make([]int, len(seedsStrs))

				for i, seed := range seedsStrs {
					seeds[i], _ = strconv.Atoi(seed)
				}

				break
			}

			continue
		}
	}

	ruleDB.Seeds = seeds
}

func populateSeedsPart2(chunks [][]string, ruleDB *RuleDB) {
	starts := []int{}
	lengths := []int{}
	seedRanges := []SeedRange{}

	for _, chunk := range chunks {
		if len(chunk) == 0 {
			continue
		}

		if len(chunk) == 1 {
			chunkStr := chunk[0]

			if seedsStr, found := strings.CutPrefix(chunkStr, "seeds: "); found {
				seedsStrs := strings.Split(seedsStr, " ")

				for i, numStr := range seedsStrs {
					num, _ := strconv.Atoi(numStr)

					if i%2 == 0 {
						starts = append(starts, num)
					} else {
						lengths = append(lengths, num)
					}
				}

				for i, start := range starts {
					seedRanges = append(seedRanges, SeedRange{SeedStart: start, SeedEnd: start + lengths[i]})
				}

				break
			}

			continue
		}
	}

	ruleDB.SeedRanges = seedRanges
}

func mapInputs(chunks [][]string, ruleDB *RuleDB) {
	seedToSoil := Rules{}
	soilToFertilizer := Rules{}
	fertilizerToWater := Rules{}
	waterToLight := Rules{}
	lightToTemperature := Rules{}
	temperatureToHumidity := Rules{}
	humidityToLocation := Rules{}

	for _, chunk := range chunks {
		if len(chunk) == 0 {
			continue
		}

		var name string
		var found bool
		if name, found = strings.CutSuffix(chunk[0], " map:"); !found {
			continue
		}

		for i := 1; i < len(chunk); i++ {
			parts := strings.Split(chunk[i], " ")
			destStart, _ := strconv.Atoi(parts[0])
			sourceStart, _ := strconv.Atoi(parts[1])
			length, _ := strconv.Atoi(parts[2])

			sourceEnd := sourceStart + length
			destOffset := destStart - sourceStart

			var target *Rules

			switch name {
			case "seed-to-soil":
				target = &seedToSoil
			case "soil-to-fertilizer":
				target = &soilToFertilizer
			case "fertilizer-to-water":
				target = &fertilizerToWater
			case "water-to-light":
				target = &waterToLight
			case "light-to-temperature":
				target = &lightToTemperature
			case "temperature-to-humidity":
				target = &temperatureToHumidity
			case "humidity-to-location":
				target = &humidityToLocation
			}

			*target = append(*target, SrcDestRange{
				SrcStart:   sourceStart,
				SrcEnd:     sourceEnd,
				DestOffest: destOffset,
			})
		}
	}

	ruleDB.SeedToSoil = seedToSoil
	ruleDB.SoilToFertilizer = soilToFertilizer
	ruleDB.FertilizerToWater = fertilizerToWater
	ruleDB.WaterToLight = waterToLight
	ruleDB.LightToTemperature = lightToTemperature
	ruleDB.TemperatureToHumidity = temperatureToHumidity
	ruleDB.HumidityToLocation = humidityToLocation
}
