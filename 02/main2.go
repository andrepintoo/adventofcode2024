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

	safeReports := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := scanner.Text()
		lvls := strings.Split(report, " ")
		intLevels := convertToIntArray(lvls)

		if isLineSafe(intLevels) {
			safeReports++
			continue
		}

		for i := 0; i < len(intLevels); i++ {
			newLine := make([]int, 0, len(intLevels)-1)
			newLine = append(newLine, intLevels[:i]...)
			newLine = append(newLine, intLevels[i+1:]...)
			if isLineSafe(newLine) {
				safeReports++
				break
			}
		}

	}

	fmt.Printf("Result: %d \n", safeReports)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func isLineSafe(levels []int) bool {
	inc, dec := true, true
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}
		if diff > 0 {
			dec = false
		} else if diff < 0 {
			inc = false
		}
	}
	return inc || dec
}

func convertToIntArray(report []string) []int {
	slice := make([]int, len(report))
	for i, s := range report {
		num, _ := strconv.Atoi(s)
		slice[i] = num
	}
	return slice
}
