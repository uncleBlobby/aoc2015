package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello, aoc-2015 day 4!")

	input, err := os.Open("input")
	if err != nil {
		log.Printf("failed to read file: %s", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(input)

	var secret string
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		secret = scanner.Text()
		// b.WriteString(scanner.Text())
	}

	var num = 1

	for {
		hasher := md5.New()
		fullString := secret + fmt.Sprintf("%d", num)
		//log.Println(fullString)
		hasher.Write([]byte(fullString))
		output := hex.EncodeToString(hasher.Sum(nil))

		if strings.HasPrefix(output, "000000") {
			log.Println("winner")
			log.Println(output)
			log.Println(num)
			break
		} else {
			num += 1
		}

	}

}
