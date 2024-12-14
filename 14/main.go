package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Robot struct{ x, y, vX, vY int }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	robots := make([]Robot, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		robots = append(robots, extractCoordinates(line))
	}

	maxX, maxY := 100, 102
	quadrants := make([]int, 4)
	midX, midY := maxX/2, maxY/2
	for _, robot := range robots {
		robot = moveRobot(robot, maxX, maxY)
		if robot.x != midX && robot.y != midY {
			i := getQuadrant(robot.x, robot.y, maxX, maxY)
			quadrants[i]++
		}
	}

	fmt.Printf("Result: %d \n", quadrants[0]*quadrants[1]*quadrants[2]*quadrants[3])
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func moveRobot(robot Robot, maxX, maxY int) Robot {
	x, y, velX, velY := robot.x, robot.y, robot.vX, robot.vY
	for i := 0; i < 100; i++ {
		x = (x + velX + (maxX + 1)) % (maxX + 1)
		y = (y + velY + (maxY + 1)) % (maxY + 1)
	}
	return Robot{x: x, y: y, vX: velX, vY: velY}
}

func getQuadrant(x, y, maxX, maxY int) int {
	if x < maxX/2 && y < maxY/2 {
		return 0
	}
	if x > maxX/2 && y < maxY/2 {
		return 1
	}
	if x < maxX/2 && y > maxY/2 {
		return 2
	}
	return 3
}

func extractCoordinates(line string) Robot {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(line, -1)

	numbers := make([]int, 4)
	for i, match := range matches {
		num, _ := strconv.Atoi(match)
		numbers[i] = num
	}

	return Robot{x: numbers[0], y: numbers[1], vX: numbers[2], vY: numbers[3]}
}
