package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	redCubes   = 12
	greenCubes = 13
	blueCubes  = 14
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("File load error: %v", err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		gameData := strings.Split(text, ":")
		rounds := strings.Split(gameData[1], ";")
		isValid := true
		for _, round := range rounds {
			sets := strings.Split(round, ",")
			for _, set := range sets {
				cubes := strings.Split(set, " ")
				amount, _ := strconv.Atoi(cubes[1])
				if cubes[2] == "red" && amount > redCubes {
					isValid = false
				}
				if cubes[2] == "green" && amount > greenCubes {
					isValid = false
				}
				if cubes[2] == "blue" && amount > blueCubes {
					isValid = false
				}
			}
		}

		if isValid {
			gameNumber, _ := strconv.Atoi(strings.Split(gameData[0], " ")[1])
			sum += gameNumber
			fmt.Println(text)
		}
	}

	fmt.Println(sum)
}
