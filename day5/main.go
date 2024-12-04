package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello, aoc-2015 day 5!")

	input, err := os.Open("input")
	if err != nil {
		log.Printf("failed to read file: %s", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(input)

	var totalNice = 0

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		totalNice += NaughtyOrNice(scanner.Text())
		// b.WriteString(scanner.Text())
	}

	log.Printf("%d strings are nice\n", totalNice)
}

func NaughtyOrNice(input string) int {
	if CheckRuleOne(input) == 0 || CheckRuleTwo(input) == 0 || CheckRuleThree(input) == 0 {
		return 0
	}
	return 1
}

func CheckRuleOne(input string) int {
	chars := strings.Split(input, "")
	totalVowels := 0
	for _, c := range chars {
		if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" {
			totalVowels += 1
		}
	}

	if totalVowels >= 3 {
		return 1
	}

	return 0
}

func CheckRuleTwo(input string) int {
	chars := strings.Split(input, "")

	for i := 0; i < len(chars)-2; i++ {
		if chars[i] == chars[i+1] {
			return 1
		}
	}

	return 0
}

func CheckRuleThree(input string) int {
	if strings.Contains(input, "ab") || strings.Contains(input, "cd") || strings.Contains(input, "pq") || strings.Contains(input, "xy") {
		return 0
	}

	return 1
}
