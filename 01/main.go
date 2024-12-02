package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	left := make([]int, 0)
	right := make([]int, 0)
	var l, r int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_, err := fmt.Sscanf(line, "%d %d", &l, &r)
		if err != nil {
			log.Fatal("error parsing line")
		}
		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	distance := 0
	for i := range left {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Printf("Result: %d \n", distance)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
