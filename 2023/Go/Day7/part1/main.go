package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Hand string
	Bid  int
}

func main() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var hands []Card

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        handed := strings.Fields(scanner.Text())
        bid, err := strconv.Atoi(handed[1])
        if err != nil {
            log.Fatal(err)
        }

        hands = append(hands, Card{
            Hand: handed[0],
            Bid: bid,
        })
    }
}

func checkPair(card Card) string {
    var cardType string
    hands := strings.Split(card.Hand, "")

    for i := 0; i < len(hands); i++ {
        if hands[i] == "" {
            continue
        }
    }

    return cardType
}


