package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func main() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	races := Race{}
	answer := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		info := raceInfo(line)

		if races.Time == 0 {
			races.Time = info
			continue
		}

		races.Distance = info
	}

	win := calcWin(races)
	if win != 0 {
		answer = answer * win
	}

	fmt.Println(answer)
}

func calcWin(race Race) int {
	count := 0
	for x := 1; x <= race.Time; x++ {
		millmeters := (race.Time - x) * x
		if millmeters > race.Distance {
			count++
		}
	}
	return count
}

func raceInfo(value string) int {
	data := strings.Split(value, ":")
	v := strings.Join(strings.Fields(data[1]), "")
	num, err := strconv.Atoi(v)
	if err != nil {
		log.Fatal(err)
	}

	return num
}
