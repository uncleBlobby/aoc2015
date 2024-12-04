package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Present struct {
	L int
	W int
	H int
}

func (p *Present) TotalSurfaceArea() int {
	return 2*p.L*p.W + 2*p.W*p.H + 2*p.H*p.L
}

func (p *Present) GetSmallestSideArea() int {
	side1 := p.L * p.W
	side2 := p.W * p.H
	side3 := p.H * p.L

	smallest := side1

	if side2 < smallest {
		smallest = side2
	}

	if side3 < smallest {
		smallest = side3
	}

	return smallest
}

func (p *Present) GetCubicVolume() int {
	return p.L * p.H * p.W
}

func (p *Present) GetBowLength() int {
	return p.GetCubicVolume()
}

func (p *Present) TotalRibbonRequired() int {
	return p.GetShortestSidePerimeter() + p.GetBowLength()
}

func (p *Present) GetShortestSidePerimeter() int {
	p1 := 2*p.L + 2*p.W
	p2 := 2*p.W + 2*p.H
	p3 := 2*p.H + 2*p.L

	smallest := p1
	if p2 < smallest {
		smallest = p2
	}
	if p3 < smallest {
		smallest = p3
	}

	return smallest
}

func (p *Present) WrappingPaper() int {
	return p.TotalSurfaceArea() + p.GetSmallestSideArea()
}

func ParsePresentFromString(measurement string) Present {
	measures := strings.Split(measurement, "x")
	ms := []int{}
	for _, measure := range measures {
		m, err := strconv.Atoi(measure)
		if err != nil {
			log.Printf("error parsing measurement: %s", err)
			continue
		}
		ms = append(ms, m)
	}

	return Present{
		L: ms[0],
		W: ms[1],
		H: ms[2],
	}
}

func main() {
	fmt.Println("hello, aoc-2015 day2!")

	input, err := os.Open("input")
	if err != nil {
		log.Printf("failed to read file: %s", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(input)

	presents := []Present{}

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		line := scanner.Text()
		p := ParsePresentFromString(line)
		presents = append(presents, p)
	}

	totalArea := 0
	totalRibbon := 0

	for _, p := range presents {
		fmt.Printf("[%d x %d x %d] = %d sqft + %d sqft\n", p.L, p.W, p.H, p.TotalSurfaceArea(), p.GetSmallestSideArea())
		totalArea += p.WrappingPaper()
		totalRibbon += p.TotalRibbonRequired()
	}

	fmt.Printf("The elves need %d sq ft of wrapping paper.\n", totalArea)
	fmt.Printf("The elves need %d feet of ribbon.\n", totalRibbon)
}
