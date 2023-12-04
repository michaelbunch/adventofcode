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
	file, err := os.Open("puzzle1data.txt")
	if err != nil {
		log.Fatalf("File load error: %v", err)
	}
	defer file.Close()

	sum := 0

	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		// Search the line for both numeric and text numbers. Don't forget about overlap numbers.
		matcher := regexp.MustCompile("[^0-9.]")
		number := ""
		for digit := range text {
			subString := text[digit:]

			if !matcher.Match([]byte(subString[:1])) {
				number = number + subString[:1]
				continue
			}

			for i, word := range words {
				if strings.Index(subString, word) == 0 {
					number = number + strconv.Itoa(i)
					continue
				}
			}
		}

		// Find the first and last digits of the string
		first := number[:1]
		last := first
		if len(number) > 1 {
			last = number[len(number)-1:]
		}

		// Make the final coordinate
		coordinates, err := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		if err != nil {
			fmt.Println("can't build coordinates")
			os.Exit(1)
		}

		// Add to the total
		sum += coordinates

		// fmt.Println(scanner.Text(), number, first, last, coordinates, sum)
	}

	fmt.Println(sum)
}
