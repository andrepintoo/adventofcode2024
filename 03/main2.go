package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type match struct {
	start, end int
	text       string
	matchType  int // 0 = dont, 1 = do, 2 = mul
}

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	r, _ := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	rDo, _ := regexp.Compile(`do\(\)`)
	rDont, _ := regexp.Compile(`don't\(\)`)
	enabled := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var matches []match
		line := scanner.Text()
		// returns [startIndex, endIndex] of each match
		exp := r.FindAllStringIndex(line, -1)
		for _, m := range exp {
			matches = append(matches, match{start: m[0], end: m[1], text: line[m[0]:m[1]], matchType: 2})
		}
		do := rDo.FindAllStringIndex(line, -1)
		fmt.Printf("do: %v\n", do)
		for _, m := range do {
			matches = append(matches, match{start: m[0], end: m[1], text: line[m[0]:m[1]], matchType: 1})
		}
		dont := rDont.FindAllStringIndex(line, -1)
		fmt.Printf("dont: %v\n", dont)
		for _, m := range dont {
			matches = append(matches, match{start: m[0], end: m[1], text: line[m[0]:m[1]], matchType: 0})
		}

		sort.Slice(matches, func(i, j int) bool {
			return matches[i].start < matches[j].start
		})

		for _, match := range matches {
			switch match.matchType {
			case 0:
				enabled = false
			case 1:
				enabled = true
			case 2:
				if enabled {
					f, s := getNumbers(match.text)
					result += (f * s)
				}
			}
		}
	}

	fmt.Printf("Result: %d \n", result)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getNumbers(reg string) (int, int) {
	// mul(23,21) -> return 23,21
	aux := reg[strings.IndexByte(reg, '('):]
	aux = strings.Trim(aux, "()")
	digits := strings.Split(aux, ",")
	f, _ := strconv.Atoi(digits[0])
	s, _ := strconv.Atoi(digits[1])

	return f, s
}
