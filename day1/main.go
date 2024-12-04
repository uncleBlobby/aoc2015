package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello, aoc-2015 day1!")

	input, err := os.Open("input")
	if err != nil {
		log.Printf("failed to read file: %s", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(input)

	var b strings.Builder

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		_, err := b.WriteString(scanner.Text())
		if err != nil {
			log.Printf("error scanning to string: %s", err)
			continue
		}
	}

	fmt.Printf("%s\n", b.String())

	chars := strings.Split(b.String(), "")

	charPos := 1
	endFloor := 0
	for _, char := range chars {
		if char == "(" {
			endFloor += 1
		}
		if char == ")" {
			endFloor -= 1
		}
		if endFloor < 0 {
			break
		}
		charPos += 1
	}

	fmt.Printf("Santa ends up on the %d floor at char position %d\n", endFloor, charPos)
}
