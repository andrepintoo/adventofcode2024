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
		safe, inc, dec := true, true, true
		lvls := strings.Split(report, " ")

		for i := 1; i < len(lvls); i++ {
			p, _ := strconv.Atoi(lvls[i-1])
			c, _ := strconv.Atoi(lvls[i])

			diff := c - p
			if diff == 0 || diff > 3 || diff < -3 {
				safe = false
				break
			}

			if diff > 0 {
				dec = false

			} else if diff < 0 {
				inc = false
			}

		}
		if safe && (inc || dec) {
			fmt.Printf("safe: %s\n", report)
			safeReports++
		}
	}

	fmt.Printf("Result: %d \n", safeReports)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
