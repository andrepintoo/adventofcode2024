package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pos struct {
	row, col int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	obstacles := map[Pos]bool{}
	guard := Pos{row: 0, col: 0}
	direction := -1 // 0 (up), 1 (right), 2 (down), 3 (left)
	maxCol := 0
	maxRow := -1
	newObstacles := []Pos{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maxRow++
		line := scanner.Text()
		for i, char := range line {
			if char == '.' {
				// every empty cell could be a new obstacule to put the guard in a loop
				newObstacles = append(newObstacles, Pos{row: maxRow, col: i})
				continue
			}
			if char == '#' {
				obstacles[Pos{row: maxRow, col: i}] = true
				continue
			}
			if direction == -1 {
				if char == '^' {
					guard = Pos{row: maxRow, col: i}
					direction = 0
				}
				if char == '>' {
					guard = Pos{row: maxRow, col: i}
					direction = 1
				}
				if char == 'v' {
					guard = Pos{row: maxRow, col: i}
					direction = 2
				}
				if char == '<' {
					guard = Pos{row: maxRow, col: i}
					direction = 3
				}
			}
		}
		if maxCol == 0 {
			// assumes all rows have the same number of columns
			maxCol = len(line)
		}

	}

	resultChan := make(chan bool, len(newObstacles))

	// now, time for some brute forcing :)
	for _, v := range newObstacles {
		// deep copy
		currentObstacles := make(map[Pos]bool)
		for k, val := range obstacles {
			currentObstacles[k] = val
		}
		currentObstacles[v] = true

		go func(curObstacles map[Pos]bool) {
			resultChan <- isGuardInLoop(guard, maxRow, maxCol, direction, curObstacles)
		}(currentObstacles)

	}

	ways := 0
	for range newObstacles {
		if <-resultChan {
			ways++
		}
	}

	fmt.Printf("Result: %d \n", ways)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// the same as part 1 logic
func isGuardInLoop(guard Pos, maxRow, maxCol, direction int, obstacles map[Pos]bool) bool {
	visitedStates := map[string]int{}
	for {
		if isGuardOut(guard, maxRow, maxCol, direction) {
			return false
		}
		stateKey := fmt.Sprintf("%d,%d,%d", guard.row, guard.col, direction)

		if visitedStates[stateKey] > 3 {
			return true
		}

		nextPos := move(guard, direction)
		if obstacles[nextPos] {
			direction = (direction + 1) % 4
			continue
		}

		guard = nextPos
		visitedStates[stateKey]++
	}
}

func move(guard Pos, direction int) Pos {
	switch direction {
	// 0 (up), 1 (right), 2 (down), 3 (left)
	case 0:
		return Pos{row: guard.row - 1, col: guard.col}
	case 1:
		return Pos{row: guard.row, col: guard.col + 1}
	case 2:
		return Pos{row: guard.row + 1, col: guard.col}
	case 3:
		return Pos{row: guard.row, col: guard.col - 1}
	}
	return guard
}

func isGuardOut(guardPos Pos, maxRow, maxCol, direction int) bool {
	if direction == 0 && guardPos.row == 0 {
		return true
	}
	if direction == 1 && guardPos.col == maxCol {
		return true
	}
	if direction == 2 && guardPos.row == maxRow {
		return true
	}
	if direction == 3 && guardPos.col == 0 {
		return true
	}
	return false
}
