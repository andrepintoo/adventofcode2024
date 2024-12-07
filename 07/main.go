package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type State struct {
	numbers []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		testValue, _ := strconv.Atoi(parts[0])
		numbersStr := strings.Fields(parts[1])

		numbers := make([]int, 0, len(numbersStr))
		for _, n := range numbersStr {
			num, _ := strconv.Atoi(n)
			numbers = append(numbers, num)
		}

		combinations := generatePossibleCombinations(numbers)
		res := calculateNumbers(numbers, combinations)
		for _, v := range res {
			if testValue == v {
				result += testValue
				break
			}
		}

	}

	fmt.Printf("Result: %d \n", result)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// ex: [2,3,4] returns ["+,+"], ["+,*"], ["*,+"], ["*,*"]
func generatePossibleCombinations(numbers []int) []string {
	operators := []string{"+", "*"}
	var combinations []string
	operatorsNeeded := len(numbers) - 1

	// recursive function
	var generate func(current string, depth int)
	generate = func(current string, depth int) {
		// base case
		if depth == operatorsNeeded {
			combinations = append(combinations, current)
			return
		}

		for _, op := range operators {
			if current == "" {
				generate(op, depth+1)
			} else {
				generate(current+","+op, depth+1)
			}
		}
	}

	generate("", 0)
	return combinations
}

func calculateNumbers(numbers []int, operations []string) []int {

	results := make([]int, 0, len(operations))

	for _, v := range operations {
		// for example, v == "+,*"
		op := strings.Split(v, ",")
		result := numbers[0]
		for j := 1; j < len(numbers); j++ {
			if op[j-1] == "+" {
				result += numbers[j]
				continue
			}
			result *= numbers[j]
		}
		results = append(results, result)
	}

	return results
}
