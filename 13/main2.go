package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct{ x, y int64 }

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
		prize := extractCoordinates(line, "=")
		prize.x += 10000000000000
		prize.y += 10000000000000
		prizes = append(prizes, prize)
	}

	cost := int64(0)
	for i, prize := range prizes {
		// thanks to cramer's rule I saw on reddit (https://www.reddit.com/r/adventofcode/comments/1hd7irq/2024_day_13_an_explanation_of_the_mathematics/)
		// A = (prize_x*b_y - prize_y*b_x) / (a_x*b_y - a_y*b_x)
		// B = (a_x*prize_y - a_y*prize_x) / (a_x*b_y - a_y*b_x)
		buttonA := buttonsA[i]
		buttonB := buttonsB[i]
		denominator := int64(buttonA.x)*int64(buttonB.y) - int64(buttonA.y)*int64(buttonB.x)
		// A
		numA := int64(prize.x)*int64(buttonB.y) - int64(prize.y)*int64(buttonB.x)

		// B
		numB := int64(buttonA.x)*int64(prize.y) - int64(buttonA.y)*int64(prize.x)

		if numA%denominator == 0 && numB%denominator == 0 {
			A := numA / denominator
			B := numB / denominator
			cost += A*3 + B
		}

	}

	fmt.Printf("Result: %d \n", cost)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func extractCoordinates(line, delimitator string) Pos {
	vals := strings.Split(line, delimitator)
	// first
	aux := strings.Split(vals[1], ",")
	unitsX := aux[0]
	valX, _ := strconv.ParseInt(unitsX, 10, 64)
	// second
	unitsY := vals[2]
	valY, _ := strconv.ParseInt(unitsY, 10, 64)

	return Pos{x: valX, y: valY}
}
