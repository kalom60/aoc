package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	answer := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		card := strings.Split(scanner.Text(), ":")[1]
		nums := strings.Split(card, "|")
		answer += getPoint(convertToNumber(nums[0]), convertToNumber(nums[1]))
	}

	fmt.Println(answer)
}

func getPoint(a, b []int) int {
	var point int = 0

	for _, num := range a {
		found := exist(num, b)
		if found {
			if point == 0 {
				point += 1
				continue
			}
			point += point
		}
	}

	return point
}

func convertToNumber(str string) []int {
	numString := strings.Split(str, " ")
	var nums []int

	for _, numStr := range numString {
		if numStr != "" {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}
	}

	return nums
}

func exist(num int, numSlice []int) bool {
	for _, number := range numSlice {
		if num == number {
			return true
		}
	}

	return false
}
