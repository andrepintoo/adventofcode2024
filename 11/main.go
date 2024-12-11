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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var stones []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stones = make([]int, 0, len(line)*2)
		nums := strings.Split(line, " ")
		for _, n := range nums {
			number, _ := strconv.Atoi(n)
			stones = append(stones, number)
		}
	}

	for i := 0; i < 25; i++ {
		var newStones []int
		for _, number := range stones {
			numStr := strconv.Itoa(number)
			if number == 0 {
				newStones = append(newStones, 1)
			} else if len(numStr)%2 == 0 {
				mid := len(numStr) / 2
				left, _ := strconv.Atoi(numStr[:mid])
				right, _ := strconv.Atoi(numStr[mid:])
				newStones = append(newStones, left, right)
			} else {
				newStones = append(newStones, number*2024)
			}
		}
		stones = newStones
	}

	fmt.Printf("Result: %d \n", len(stones))
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
