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

	// _ = NaughtyOrNice("ugknbfddgicrmopn")
	// _ = NaughtyOrNice("aaa")
	// _ = NaughtyOrNice("jchzalrnumimnmhp")
	// _ = NaughtyOrNice("hhaegwjzuvuyypxyu")
	// _ = NaughtyOrNice("dvszwmarrgswjxmb")

	// _ = NaughtyOrNice2("qjhvhtzxzqqjkmpb")
	// _ = NaughtyOrNice2("xxyxx")
	// _ = NaughtyOrNice2("uurcxstgmygtbstg")
	// _ = NaughtyOrNice2("ieodomkazucvgmuy")

	scanner := bufio.NewScanner(input)

	var totalNice = 0

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		totalNice += NaughtyOrNice2(scanner.Text())
		// b.WriteString(scanner.Text())
	}

	log.Printf("%d strings are nice\n", totalNice)
}

func NaughtyOrNice(input string) int {
	if CheckRuleOne(input) == 0 || CheckRuleTwo(input) == 0 || CheckRuleThree(input) == 0 {
		fmt.Printf("%s is naughty\n", input)
		return 0
	}
	fmt.Printf("%s is nice\n", input)
	return 1
}

func NaughtyOrNice2(input string) int {
	if NewRuleOne(input) == 0 || NewRuleTwo(input) == 0 {
		fmt.Printf("%s is naughty\n", input)
		return 0
	}
	fmt.Printf("%s is nice\n", input)
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

	if chars[len(chars)-1] == chars[len(chars)-2] {
		return 1
	}

	return 0
}

func CheckRuleThree(input string) int {
	if strings.Contains(input, "ab") || strings.Contains(input, "cd") || strings.Contains(input, "pq") || strings.Contains(input, "xy") {
		return 0
	}

	return 1
}

func NewRuleOne(input string) int {

	chars := strings.Split(input, "")

	pairs := make([][]string, 0)

	for i := 0; i < len(chars)-1; i += 1 {
		newPair := []string{chars[i], chars[i+1]}
		pairs = append(pairs, newPair)
	}

	for i := 0; i < len(pairs); i++ {
		for j := 1; j < len(pairs); j++ {
			if pairs[i][0] == pairs[j][0] && pairs[i][1] == pairs[j][1] {
				if i != j {
					return 1
				}
			}
		}
	}
	//fmt.Printf("%s failed rule1", input)
	return 0
}

func NewRuleTwo(input string) int {
	//contains at least one letter which repeats with exactly one letter between them

	chars := strings.Split(input, "")

	for i := 0; i+2 < len(chars); i++ {
		if chars[i] == chars[i+2] {
			return 1
		}
	}
	// fmt.Printf("%s failed rule2", input)
	return 0
}
