package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func main(){
	file, err := os.Open("input.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
	}

	fmt.Printf("Result: %d \n", ..)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
}
