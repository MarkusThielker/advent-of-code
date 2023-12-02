package main

import (
	"fmt"
	utils "github.com/markusthielker/advent-of-code/2023/go"
	"strconv"
	"unicode"
)

func main() {

	input := utils.ReadInput("day-01/input")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) (result int) {

	for _, line := range input {

		firstDigit, ok1 := findFirst(line, unicode.IsDigit)
		lastDigit, ok2 := findLast(line, unicode.IsDigit)

		if ok1 && ok2 {

			firstDigitTyped, _ := strconv.Atoi(firstDigit)
			lastDigitTyped, _ := strconv.Atoi(lastDigit)

			result += firstDigitTyped*10 + lastDigitTyped
		}
	}

	return
}

func part2(input []string) (result int) {
	return len(input)
}

func findFirst(input string, condition func(rune) bool) (string, bool) {

	for _, char := range input {
		if condition(char) {
			return string(char), true
		}
	}

	return "", false
}

func findLast(input string, condition func(rune) bool) (string, bool) {

	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		if condition(runes[i]) {
			return string(runes[i]), true
		}
	}

	return "", false
}
