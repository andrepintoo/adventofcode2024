package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
		if !isRightOrder(nums, rules) {
			correctOrder := []int{}
			for _, page := range nums {
				pos := findPosition(correctOrder, page, rules)
				// insert in the correct position
				correctOrder = append(correctOrder[:pos], append([]int{page}, correctOrder[pos:]...)...)
			}
			result += correctOrder[len(correctOrder)/2]
		}
	}

	fmt.Printf("Result: %d \n", result)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// part 1 logic
func isRightOrder(nums []int, rules map[int][]int) bool {

	indexes := make(map[int]int)
	for idx, num := range nums {
		indexes[num] = idx
	}
	for _, num := range nums {
		// check the rules of this specific page number in the update
		for _, afterNum := range rules[num] {
			if _, exists := indexes[afterNum]; !exists {
				continue
			}
			if indexes[afterNum] < indexes[num] {
				return false
			}
		}
	}

	return true
}

// check if the current page should come before a page in "correct order", then return that existing page index
func findPosition(correctOrder []int, currentPage int, rules map[int][]int) int {
	for idx, existingPage := range correctOrder {
		if slices.Contains(rules[currentPage], existingPage) {
			return idx
		}
	}
	return len(correctOrder)
}

func getPairNumbers(line string) (int, int) {
	// 12|3 --> returns (12,3)
	parts := strings.Split(line, "|")
	num1, _ := strconv.Atoi(parts[0])
	num2, _ := strconv.Atoi(parts[1])
	return num1, num2
}
