package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var contents [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, []rune(scanner.Text()))
	}

	var answer int = 0

	for y := 0; y < len(contents); y++ {
		row := contents[y]

		for x := 0; x < len(row); x++ {
			if isStar(contents[y][x]) {
				pr, nr := y-1, y+1
				if 0 <= pr && pr < len(contents) && 0 <= nr && nr <= len(contents) {
					partNumbers := getPartNumbers(1, x, contents[pr:nr+1])
					if len(partNumbers) > 1 {
						answer += partNumbers[0] * partNumbers[1]
					}
				}
			}
		}
	}

	fmt.Println(answer)
}

func getPartNumbers(i, z int, schematic [][]rune) []int {
	var partNumbers []int
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for y := 0; y < len(schematic); y++ {
		row := schematic[y]
		start := 0
		j := 0

		for j < len(row) {
			currNumber := ""
			start = j

			for j < len(row) && isCharNumber(row[j]) {
				currNumber += string(row[j])
				j++
			}

			if currNumber != "" {
				num, _ := strconv.Atoi(currNumber)
				foundSymbol := false

				for _, dir := range directions {
					nr, nc := y+dir[0], j+dir[1]-1
					if (0 <= nr && nr < len(schematic)) && (0 <= nc && nc < len(row)) {
						if nr == i && nc == z {
							partNumbers = append(partNumbers, num)
							foundSymbol = true
							break
						}
					}
				}

				if !foundSymbol {
					for _, dir := range directions {
						nr, nc := y+dir[0], start+dir[1]
						if (0 <= nr && nr < len(schematic)) && (0 <= nc && nc < len(row)) {
							if nr == i && nc == z {
								partNumbers = append(partNumbers, num)
								break
							}
						}
					}
				}
			}
			j++
		}
	}

	return partNumbers
}

func isCharNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isStar(r rune) bool {
	return r == '*'
}
