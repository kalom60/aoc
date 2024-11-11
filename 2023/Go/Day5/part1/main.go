package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var maps = [7]string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light",
	"light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

func main() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var seeds []int
	// var lowestLocation int
	mapping := make(map[string][][]int)
	currentMapping := ""

	validMappings := make(map[string]bool)
	for _, m := range maps {
		validMappings[m] = true
	}

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		seedValues := strings.Split(line, ":")
		if len(seedValues) > 1 {
			seedNum := strings.Fields(strings.TrimSpace(seedValues[1]))
			seeds = append(seeds, convertToNumber(seedNum)...)
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		mapper := strings.Split(line, " ")[0]

		if mapper != "" {
			if _, found := validMappings[mapper]; found {
				currentMapping = mapper

				if _, exist := mapping[currentMapping]; !exist {
					mapping[currentMapping] = [][]int{}
				}
				continue
			}

			parts := strings.Fields(line)
			mapping[currentMapping] = append(mapping[currentMapping], convertToNumber(parts))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, m := range maps {
		ranges := mapping[m]
		fmt.Println(m, ranges)

        for _, seed := range seeds {
            fmt.Printf("%d\t", seed)
        }
        fmt.Println()
	}
}

func convertToNumber(input []string) []int {
	var nums []int

	for _, v := range input {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, num)
	}

	return nums
}
