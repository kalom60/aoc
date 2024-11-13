package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type RangeMapping struct {
	destStart, sourceStart, length int
}

func main() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maps := make(map[string][]RangeMapping, 7)
	maps["seed-to-soil"] = []RangeMapping{}
	maps["soil-to-fertilizer"] = []RangeMapping{}
	maps["fertilizer-to-water"] = []RangeMapping{}
	maps["water-to-light"] = []RangeMapping{}
	maps["light-to-temperature"] = []RangeMapping{}
	maps["temperature-to-humidity"] = []RangeMapping{}
	maps["humidity-to-location"] = []RangeMapping{}

	var seeds string
	currentMapping := ""

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		seeds = strings.Split(line, ":")[1]
	}

	seedRanges := parseSeedRanges(seeds)

	for scanner.Scan() {
		line := scanner.Text()
		mapper := strings.Split(line, " ")[0]

		if mapper != "" {
			if _, found := maps[mapper]; found {
				currentMapping = mapper
				continue
			}

			parts := strings.Fields(line)
			maps[currentMapping] = append(maps[currentMapping], NewRangeMapping(parts))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	mapRange := func(ranges []RangeMapping, num int) int {
		for _, r := range ranges {
			if num >= r.sourceStart && num < r.sourceStart+r.length {
				return r.destStart + (num - r.sourceStart)
			}
		}
		return num
	}

	minLoc := math.MaxInt
	for _, start := range seedRanges {
		for i := 0; i < start.length; i++ {
			seed := start.start + i

			soil := mapRange(maps["seed-to-soil"], seed)
			fertilizer := mapRange(maps["soil-to-fertilizer"], soil)
			water := mapRange(maps["fertilizer-to-water"], fertilizer)
			light := mapRange(maps["water-to-light"], water)
			temperature := mapRange(maps["light-to-temperature"], light)
			humidity := mapRange(maps["temperature-to-humidity"], temperature)
			location := mapRange(maps["humidity-to-location"], humidity)

			if location < minLoc {
				minLoc = location
			}
		}
	}

	fmt.Println(minLoc)
}

func parseSeedRanges(input string) []struct{ start, length int } {
	parts := strings.Fields(input)
	var ranges []struct{ start, length int }
	for i := 0; i < len(parts); i += 2 {
		start := atoi(parts[i])
		length := atoi(parts[i+1])
		ranges = append(ranges, struct{ start, length int }{start, length})
	}
	return ranges
}

func atoi(s string) int {
	var res int
	for _, c := range s {
		res = res*10 + int(c-'0')
	}
	return res
}

func NewRangeMapping(input []string) RangeMapping {
	var nums []int

	for _, v := range input {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, num)
	}

	return RangeMapping{nums[0], nums[1], nums[2]}
}
