package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var total, total2 int

var cubes = map[string]int{
	"blue":  14,
	"green": 13,
	"red":   12,
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	stringData := strings.Split(string(data), "\n")
	for _, str := range stringData {
		valid, id := isValidGame(str)
		if valid {
			numId, _ := strconv.Atoi(id)
			total += numId
		}
		total2 += setPower(str)
	}
	fmt.Printf("TOTAL: %v\n", total)
	fmt.Printf("TOTAL 2: %v\n", total2)
}

// PROBLEM 1

func isValidGame(game string) (bool, string) {
	splittedLine := strings.Split(game, ":")
	gameId := strings.Split(splittedLine[0], " ")[1]
	cubeDraws := strings.Split(splittedLine[1], ";")
	for _, v := range cubeDraws {
		if !isValidDraw(v) {
			return false, gameId
		}
	}
	return true, gameId

}

func isValidDraw(draw string) bool {

	splittedDraw := strings.Split(draw, ",")
	for _, v := range splittedDraw {
		spl := strings.Split(v, " ")
		count, _ := strconv.Atoi(spl[1])
		color := spl[2]
		if cubes[color] < count {
			return false
		}
	}
	return true
}

// PROBLEM 2

func setPower(set string) int {

	colorMap := make(map[string]int)
	splittedLine := strings.Split(set, ":")
	cubeDraws := strings.Split(splittedLine[1], ";")
	for _, v := range cubeDraws {
		splittedDraw := strings.Split(v, ",")
		for _, v := range splittedDraw {
			spl := strings.Split(v, " ")
			count, _ := strconv.Atoi(spl[1])
			color := spl[2]
			if count > colorMap[color] {
				colorMap[color] = count
			}
		}
	}

	return colorMap["green"] * colorMap["blue"] * colorMap["red"]
}
