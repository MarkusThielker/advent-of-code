package main

import (
	"fmt"
	utils "github.com/markusthielker/advent-of-code/2023/go"
	"strconv"
	"strings"
)

func main() {

	part1 := func(input []string) (result int) {

		for _, card := range input {
			parts := strings.Split(card, "|")
			winningNumbers := toIntSlice(strings.Fields(strings.Split(parts[0], ":")[1]))
			cardNumbers := toIntSlice(strings.Fields(parts[1]))

			wins := 0
			for _, num := range cardNumbers {
				if contains(winningNumbers, num) {
					wins++
				}
			}

			points := 0
			if wins > 0 {
				points = 1
				for i := 1; i < wins; i++ {
					points *= 2
				}
			}

			result += points
		}

		return
	}

	part2 := func(input []string) (result int) {
		return len(input)
	}

	input := utils.ReadInput("day-04/input")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func toIntSlice(strings []string) []int {
	var ints []int
	for _, s := range strings {
		if s != "" {
			if num, err := strconv.Atoi(s); err == nil {
				ints = append(ints, num)
			}
		}
	}
	return ints
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
