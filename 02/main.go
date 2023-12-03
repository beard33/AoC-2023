package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var total int

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
			fmt.Printf("ID: %v\n\n", id)
			total += numId
		}
	}
	fmt.Printf("TOTAL: %v\n", total)
}

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
			fmt.Printf("Count: %v Color: %v\n", count, color)
			fmt.Println("Impossible!")
			return false
		}
	}

	return true

}
