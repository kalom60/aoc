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

	var races []Race
    answer := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
        info := raceInfo(line)

        for i, r := range info {
            if len(races) < len(info) {
                race := Race{
                    Time: r,
                }

                races = append(races, race)
                continue
            }

           races[i].Distance = r
        }
	}

    for _, r := range races {
        win := calcWin(r)
        if win != 0 {
            answer = answer * win
        }
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

func raceInfo(value string) []int {
	data := strings.Split(value, ":")
	return convertToNumber(strings.Fields(data[1]))
}

func convertToNumber(str []string) []int {
	var nums []int

	for _, s := range str {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, num)
	}

	return nums
}
