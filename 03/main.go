package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	r, _ := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		exp := r.FindAllString(line, -1)
		for _, match := range exp {
			f, s := getNumbers(match)
			result += (f * s)
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
