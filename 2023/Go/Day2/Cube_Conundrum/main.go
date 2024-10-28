package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var cardInfo = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var countID int = 0
	var fewestCards int = 0

	for scanner.Scan() {
		var cards = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		var fewestCubs = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		var flag bool = true

		line := scanner.Text()

		list := strings.Split(line, ":")
		id := strings.Split(list[0], " ")[1]
		list[1] = strings.TrimLeft(list[1], " ")
		games := strings.Split(list[1], ";")

		for _, game := range games {
			cardInput := strings.Split(game, ",")
			for _, card := range cardInput {
				cardItems := strings.Split(strings.TrimLeft(card, " "), " ")
				cardValue, err := strconv.Atoi(cardItems[0])
				if err != nil {
					log.Fatal(err)
				}

				if cardValue >= fewestCubs[cardItems[1]] {
					fewestCubs[cardItems[1]] = cardValue
				}

				if ok := cardValue <= cardInfo[cardItems[1]]; ok && flag {
					cards[cardItems[1]] += cardValue
					flag = true
				} else {
					flag = false
				}
			}
		}

		fewestCards += fewestCubs["red"] * fewestCubs["green"] * fewestCubs["blue"]

		if flag {
			gameID, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}

			countID += gameID
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Count ID and Fewest Cubs need is %d, %d\n ", countID, fewestCards)
}
