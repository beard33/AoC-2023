package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var total1, total2 int
var numMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"ten":   "10",
}

func main() {

	data, _ := os.ReadFile("./input.txt")
	stringData := strings.Split(string(data), "\n")

	for _, str := range stringData {
		strNumber := findNumbers(str)
		number, _ := strconv.Atoi(strNumber)
		total1 += number

		n1, n2 := findLiteralsNumber(str)
		toBeConverted := fmt.Sprintf("%v%v", n1, n2)
		newNumber, _ := strconv.Atoi(toBeConverted)
		fmt.Println(newNumber)
		total2 += newNumber
	}
	fmt.Printf("Total for problem 1: %v\n", total1)
	fmt.Printf("Total for problem 2: %v\n", total2)

}

func findNumbers(str string) (out string) {
	var n1, n2 string

	for _, c := range str {
		if unicode.IsNumber(c) {
			if n1 == "" {
				n1 = string(c)
				continue
			}
			n2 = string(c)
		}
	}

	out = fmt.Sprintf("%v%v", n1, n2)
	if n2 == "" {
		out = fmt.Sprintf("%v%v", n1, n1)
	}
	return
}

func findLiteralsNumber(str string) (n1, n2 string) {
	var i1, i3 = 500, 500
	var i2, i4 = -1, -1
	var n3, n4 string

	for k := range numMap {
		numFirstIndex := strings.Index(str, k)
		numLastIndex := strings.LastIndex(str, k)

		if numFirstIndex >= 0 && numFirstIndex < i1 {
			i1 = numFirstIndex
			n1 = numMap[k]
		}

		if numLastIndex > i2 {
			i2 = numLastIndex
			n2 = numMap[k]
		}
	}

	for i, c := range str {
		if unicode.IsNumber(c) {
			if n3 == "" {
				n3 = string(c)
				i3 = i
			}
			n4 = string(c)
			i4 = i
		}
	}

	if i3 < i1 {
		n1 = n3
	}

	if i4 > i2 {
		n2 = n4
	}
	return
}
