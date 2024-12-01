package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	left := make([]int, 0)
	right := make(map[int]int)
	var l, r int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_, err := fmt.Sscanf(line, "%d %d", &l, &r)
		if err != nil {
			log.Fatal("error parsing line")
		}
		left = append(left, l)
		right[r]++
	}

	score := 0
	for _, v := range left {
		occurrences := right[v]
		score += (v * occurrences)
	}

	fmt.Printf("Result: %d \n", score)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
