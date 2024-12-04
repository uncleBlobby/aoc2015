package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Position struct {
	X int
	Y int
}

type Santa struct {
	Position Position
}

func (s *Santa) Move(instruction string) {
	switch instruction {
	case ">":
		s.Position.X += 1
	case "<":
		s.Position.X -= 1
	case "^":
		s.Position.Y += 1
	case "v":
		s.Position.Y -= 1
	}
}

func (s *Santa) DeliverGift(houses map[Position]int) {
	houses[s.Position] += 1
}

func main() {
	fmt.Println("hello, aoc-2015 day3!")

	input, err := os.Open("input")
	if err != nil {
		log.Printf("failed to read file: %s", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(input)

	var b strings.Builder
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		b.WriteString(scanner.Text())
	}

	instructions := strings.Split(b.String(), "")

	s := Santa{
		Position: Position{X: 0, Y: 0},
	}

	rs := Santa{
		Position: Position{X: 0, Y: 0},
	}

	houses := map[Position]int{}
	robohouses := map[Position]int{}
	s.DeliverGift(houses)
	rs.DeliverGift(robohouses)

	fmt.Printf("%d total instructions\n", len(instructions))
	for ind, inst := range instructions {
		//fmt.Printf("%s", inst)
		if ind == 0 {
			s.Move(inst)
			s.DeliverGift(houses)
		} else if ind == 1 {
			rs.Move(inst)
			rs.DeliverGift(robohouses)
		} else if ind%2 == 0 {
			s.Move(inst)
			s.DeliverGift(houses)
		} else if ind%2 != 0 {
			rs.Move(inst)
			rs.DeliverGift(robohouses)
		}
	}

	// for position, presents := range houses {
	// 	//log.Printf("%v\t%d", position, presents)
	// }

	fmt.Printf("%d unique houses\n", len(houses)+len(robohouses))

	totalHouses := map[Position]int{}

	for santahouse := range houses {
		if totalHouses[santahouse] == 0 {
			totalHouses[santahouse] += 1
		}
	}

	for robohouse := range robohouses {
		if totalHouses[robohouse] == 0 {
			totalHouses[robohouse] += 1
		}
	}

	fmt.Printf("%d total houses", len(totalHouses))
}
