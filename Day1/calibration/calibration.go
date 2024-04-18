package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var textNumbers = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	file, err := os.Open("calibration.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count int

	for scanner.Scan() {
		var firstDigit int = -1
		var lastDigit int = -1

		line := scanner.Text()

        replaceTextWithNumber(&line)

		findFirstAndLastDigit(line, &firstDigit, &lastDigit)

		concatAndIncrementCount(&firstDigit, &lastDigit, &count)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The sum of my calibration value is %d", count)
}

func replaceTextWithNumber(line *string) {
	for pos, textNum := range textNumbers {
		if ok := strings.Contains(*line, textNum); ok {
			*line = strings.ReplaceAll(*line, textNum, fmt.Sprintf("%s%d%s", string(textNum[0]), pos+1, string(textNum[len(textNum)-1])))
		}
	}
}

func findFirstAndLastDigit(line string, firstDigit *int, lastDigit *int) {
	for _, char := range line {
		if unicode.IsDigit(char) {
			if *firstDigit == -1 {
				*firstDigit = int(char - '0')
			} else {
				*lastDigit = int(char - '0')
			}
		}
	}
}

func concatAndIncrementCount(firstDigit, lastDigit, count *int) {
	if *firstDigit != -1 && *firstDigit > 0 && *lastDigit > -1 {
		numberStr := fmt.Sprintf("%d%d", *firstDigit, *lastDigit)

		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatal(err)
		}

		*count += number
	} else if *firstDigit != -1 && *firstDigit > 0 && *lastDigit == -1 {
		numberStr := fmt.Sprintf("%d%d", *firstDigit, *firstDigit)

		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatal(err)
		}

		*count += number
	}
}
