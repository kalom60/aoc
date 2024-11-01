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

	var answer int
	instance := make(map[int][][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		table := strings.Split(scanner.Text(), ":")
		card := strings.Split(table[0], " ")
		cardNumber, _ := strconv.Atoi(card[len(card)-1])
		nums := strings.Split(table[1], "|")
		matching := getPoint(convertToNumber(nums[0]), convertToNumber(nums[1]), cardNumber)

		instance[cardNumber] = [][]int{matching}
		if len(matching) > 0 {
			cardCopy := foundCopy(instance, cardNumber)
			if cardCopy > 0 {
				for i := 1; i <= cardCopy; i++ {
					instance[cardNumber] = append(instance[cardNumber], matching)
				}
			}
		}
	}

	for _, value := range instance {
		answer++
		for _, val := range value {
			for range val {
				answer++
			}
		}
	}

	fmt.Println(answer)
}

func getPoint(a, b []int, start int) []int {
	var point = []int{}

	for _, num := range a {
		found := exist(num, b)
		if found {
			if len(point) == 0 {
				point = append(point, start+1)
				continue
			}

			lastEle := len(point) - 1
			point = append(point, point[lastEle]+1)
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

func foundCopy(instance map[int][][]int, card int) int {
	var copied int = 0

	for i := 1; i <= len(instance); i++ {
		if value, found := instance[card-i]; found {
			if numExist := exist(card, value[0]); numExist {
				copied += len(value)
			}
		}
	}

	return copied
}
