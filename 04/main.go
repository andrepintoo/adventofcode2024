package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	matrix := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}

	xmas := []string{"X", "M", "A", "S"}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == xmas[0] {
				// XMAS
				if matchPattern(matrix, i, j, 0, 1, xmas) {
					result++
				}

				// SAMX
				if matchPattern(matrix, i, j, 0, -1, xmas) {
					result++
				}

				// vertical XMAS downwards
				if matchPattern(matrix, i, j, 1, 0, xmas) {
					result++
				}

				// vertical XMAS upwards
				if matchPattern(matrix, i, j, -1, 0, xmas) {
					result++
				}

				//X
				// M
				//  A
				//   S
				if matchPattern(matrix, i, j, 1, 1, xmas) {
					result++
				}

				//   S
				//  A
				// M
				//X
				if matchPattern(matrix, i, j, -1, -1, xmas) {
					result++
				}

				//S
				// A
				//  M
				//   X
				if matchPattern(matrix, i, j, 1, -1, xmas) {
					result++
				}

				//   S
				//  A
				// M
				//X
				if matchPattern(matrix, i, j, -1, 1, xmas) {
					result++
				}
			}
		}
	}

	fmt.Printf("Result: %d\n", result)

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// pattern gotten after some brute forcing :)
func matchPattern(matrix [][]string, startRow, startCol, rowStep, colStep int, pattern []string) bool {
	for k := 0; k < len(pattern); k++ {
		r := startRow + k*rowStep
		c := startCol + k*colStep

		if r < 0 || r >= len(matrix) || c < 0 || c >= len(matrix[0]) {
			return false
		}

		if matrix[r][c] != pattern[k] {
			return false
		}
	}
	return true
}
