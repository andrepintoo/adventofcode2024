package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Pos struct{ row, col int }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := [][]int{}
	trailheads := []Pos{}
	scanner := bufio.NewScanner(file)
	maxRow := 0
	for scanner.Scan() {
		line := scanner.Text()
		l := make([]int, 0, len(line))
		for j, c := range line {
			n, _ := strconv.Atoi(string(c))
			if n == 0 {
				trailheads = append(trailheads, Pos{row: maxRow, col: j})
			}
			l = append(l, n)
		}
		m = append(m, l)
		maxRow++
	}

	// each move can potentially lead to another 4 moves
	totalScore := 0
	for _, j := range trailheads {
		//make sure to not repeat heads
		visitedNines := make(map[Pos]bool)
		totalScore += searchNines(m, j, 0, visitedNines)
	}

	fmt.Printf("Result: %d \n", totalScore)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// returns the score from that trailhead
func searchNines(matrix [][]int, pos Pos, currentHeight int, nines map[Pos]bool) int {
	rows := len(matrix)
	cols := len(matrix[0])

	if matrix[pos.row][pos.col] == 9 {
		if !nines[pos] {
			nines[pos] = true
			return 1
		}
		// only unique nines are counted
		return 0
	}

	score := 0
	// can go to the right, left, top or bottom
	directions := []Pos{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for _, dir := range directions {
		nextPos := Pos{row: pos.row + dir.row, col: pos.col + dir.col}
		if nextPos.row < 0 || nextPos.col < 0 || nextPos.row >= rows || nextPos.col >= cols {
			continue
		}
		if matrix[nextPos.row][nextPos.col] == (currentHeight + 1) {
			score += searchNines(matrix, nextPos, currentHeight+1, nines)
		}
	}
	return score
}
