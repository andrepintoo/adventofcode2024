package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var memo = make(map[string]int)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var stones []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		for _, n := range nums {
			number, _ := strconv.Atoi(n)
			stones = append(stones, number)
		}
	}

	totalStones := 0
	// process each stone individually
	for _, stone := range stones {
		totalStones += processStone(stone, 75)
	}

	fmt.Printf("Result: %d\n", totalStones)

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processStone(stone, blinks int) int {
	// base case
	if blinks == 0 {
		return 1
	}

	key := fmt.Sprintf("%d:%d", stone, blinks)

	if result, found := memo[key]; found {
		return result
	}

	var result int
	if stone == 0 {
		result = processStone(1, blinks-1)
	} else {
		numStr := strconv.Itoa(stone)
		if len(numStr)%2 == 0 {
			mid := len(numStr) / 2
			left, _ := strconv.Atoi(numStr[:mid])
			right, _ := strconv.Atoi(numStr[mid:])
			result = processStone(left, blinks-1) + processStone(right, blinks-1)
		} else {
			result = processStone(stone*2024, blinks-1)
		}
	}

	memo[key] = result

	return result
}
