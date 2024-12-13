package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct{ x, y int }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	prizes := []Pos{}
	buttonsA := []Pos{}
	buttonsB := []Pos{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "Button A") {
			buttonsA = append(buttonsA, extractCoordinates(line, "+"))
			continue
		}
		if strings.Contains(line, "Button B") {
			buttonsB = append(buttonsB, extractCoordinates(line, "+"))
			continue
		}
		prizes = append(prizes, extractCoordinates(line, "="))
	}

	count := 0
	for i, v := range prizes {
		reachable, aCount, bCount := breathFirstSearchLookup(v, buttonsA[i], buttonsB[i])
		if reachable {
			count += (aCount*3 + bCount)
		}
	}

	fmt.Printf("Result: %d \n", count)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func breathFirstSearchLookup(prize, buttonA, buttonB Pos) (bool, int, int) {
	type State struct {
		pos    Pos
		aCount int
		bCount int
	}

	queue := []State{{pos: Pos{0, 0}, aCount: 0, bCount: 0}}
	// some memoization or otherwise program will not complete
	visited := make(map[string]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// base case
		if current.pos.x == prize.x && current.pos.y == prize.y {
			// fmt.Printf("Found prize with A presses: %d, B presses: %d\n", current.aCount, current.bCount)
			return true, current.aCount, current.bCount
		}

		if current.aCount > 100 || current.bCount > 100 {
			continue
		}

		key := fmt.Sprintf("%d,%d,%d,%d", current.pos.x, current.pos.y, current.aCount, current.bCount)
		if visited[key] {
			continue
		}
		visited[key] = true

		// press A
		queue = append(queue, State{
			pos:    Pos{x: current.pos.x + buttonA.x, y: current.pos.y + buttonA.y},
			aCount: current.aCount + 1,
			bCount: current.bCount,
		})

		// press B
		queue = append(queue, State{
			pos:    Pos{x: current.pos.x + buttonB.x, y: current.pos.y + buttonB.y},
			aCount: current.aCount,
			bCount: current.bCount + 1,
		})
	}

	return false, 0, 0
}

func extractCoordinates(line, delimitator string) Pos {
	vals := strings.Split(line, delimitator)
	// first
	aux := strings.Split(vals[1], ",")
	unitsX := aux[0]
	valX, _ := strconv.Atoi(unitsX)
	// second
	unitsY := vals[2]
	valY, _ := strconv.Atoi(unitsY)

	return Pos{x: valX, y: valY}
}
