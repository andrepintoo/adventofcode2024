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

	searchXmas(matrix, &result)

	fmt.Printf("Result: %d \n", result)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func searchXmas(m [][]string, res *int) {
	rows := len(m)
	cols := len(m[0])
	mas := []string{"M", "A", "S"}
	sam := []string{"S", "A", "M"}

	// check by the center
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if isValidCenter(m, r, c, mas, mas) ||
				isValidCenter(m, r, c, mas, sam) ||
				isValidCenter(m, r, c, sam, mas) ||
				isValidCenter(m, r, c, sam, sam) {
				*res++
			}
		}
	}
}

func isValidCenter(matrix [][]string, r, c int, topLeftToBottomRight, topRightToBottomLeft []string) bool {
	// this will always be inside the boundaries because of restrictions in the nested for loop in searchDiagonal()
	return matrix[r][c] == "A" &&
		matrix[r-1][c-1] == topLeftToBottomRight[0] && matrix[r+1][c+1] == topLeftToBottomRight[2] &&
		matrix[r-1][c+1] == topRightToBottomLeft[2] && matrix[r+1][c-1] == topRightToBottomLeft[0]
}
