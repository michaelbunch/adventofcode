package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("puzzle1data.txt")
	if err != nil {
		log.Fatalf("File load error: %v", err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matcher := regexp.MustCompile("[^0-9.]")
		onlyDigits := string(matcher.ReplaceAll([]byte(scanner.Text()), []byte("")))

		first := onlyDigits[:1]
		last := first
		if len(onlyDigits) > 1 {
			last = onlyDigits[len(onlyDigits)-1:]
		}

		coordinates, _ := strconv.Atoi(first + last)
		sum += coordinates
	}

	fmt.Println(sum)
}
