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

	var sum int = 0
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for y := 0; y < len(contents); y++ {
		row := contents[y]
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
					if (0 <= nr && nr < len(contents)) && (0 <= nc && nc < len(row)) {
						if isSymbol(contents[nr][nc]) {
							sum += num
							foundSymbol = true
							break
						}
					}
				}

				if !foundSymbol {
					for _, dir := range directions {
						nr, nc := y+dir[0], start+dir[1]
						if (0 <= nr && nr < len(contents)) && (0 <= nc && nc < len(row)) {
							if isSymbol(contents[nr][nc]) {
								sum += num
								break
							}
						}
					}
				}
			}
			j++
		}
	}

	fmt.Println(sum)
}

func isCharNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isSymbol(r rune) bool {
	return r != '.' && !isCharNumber(r)
}
