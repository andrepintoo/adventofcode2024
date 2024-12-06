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

	rules := make(map[int][]int)
	scanner := bufio.NewScanner(file)
	updates := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			n1, n2 := getPairNumbers(line)
			rules[n1] = append(rules[n1], n2)
			continue
		}
		if len(line) == 0 {
			continue
		}
		nums := strings.Split(line, ",")
		update := []int{}
		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			update = append(update, n)
		}
		updates = append(updates, update)
	}

	result := 0
	for _, nums := range updates {
		rightOrder := true

		// stores the order of the page numbers in this update
		indexes := make(map[int]int)
		for idx, num := range nums {
			indexes[num] = idx
		}
		for _, num := range nums {
			// check the rule (numbers that num has to precede) for each page number in this update
			for _, rule := range rules[num] {
				// if a number in its rule isn't present in the update, skip
				if _, exists := indexes[rule]; !exists {
					continue
				}
				// when a page that should go after this number precedes it, then the order is wrong
				if indexes[rule] < indexes[num] {
					// fmt.Printf(" incorrect in %v: %d must be after %d\n", nums, rule, num)
					rightOrder = false
					break
				}
			}
			if !rightOrder {
				// no need to continue checking in case of failure
				break
			}
		}
		if rightOrder {
			result += nums[len(nums)/2]
		}
	}

	fmt.Printf("Result: %d \n", result)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func getPairNumbers(line string) (int, int) {
	// 12|3 --> returns (12,3)
	parts := strings.Split(line, "|")
	num1, _ := strconv.Atoi(parts[0])
	num2, _ := strconv.Atoi(parts[1])
	return num1, num2
}
