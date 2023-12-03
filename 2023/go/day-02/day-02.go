package main

import (
	"fmt"
	utils "github.com/markusthielker/advent-of-code/2023/go"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	nonDigitRegex := regexp.MustCompile(`[^0-9]`)

	part1 := func(input []string) (result int) {

		result, redLimit, greenLimit, blueLimit := 0, 12, 13, 14

	game:
		for index, game := range input {

			grabs := strings.Split(strings.Split(game, ":")[1], ";")
			for _, grab := range grabs {

				colors := strings.Split(grab, ",")
				for _, color := range colors {

					digit := nonDigitRegex.ReplaceAllString(color, "")
					digitTyped, _ := strconv.Atoi(digit)

					if strings.Contains(color, "red") && digitTyped > redLimit {
						continue game
					} else if strings.Contains(color, "green") && digitTyped > greenLimit {
						continue game
					} else if strings.Contains(color, "blue") && digitTyped > blueLimit {
						continue game
					}
				}
			}

			result += index + 1
		}

		return
	}

	part2 := func(input []string) (result int) {

		result = 0

		for _, game := range input {

			redMax, greenMax, blueMax := 0, 0, 0

			grabs := strings.Split(strings.Split(game, ":")[1], ";")
			for _, grab := range grabs {

				colors := strings.Split(grab, ",")
				for _, color := range colors {

					digit := nonDigitRegex.ReplaceAllString(color, "")
					digitTyped, _ := strconv.Atoi(digit)

					if strings.Contains(color, "red") && digitTyped > redMax {
						redMax = digitTyped
					} else if strings.Contains(color, "green") && digitTyped > greenMax {
						greenMax = digitTyped
					} else if strings.Contains(color, "blue") && digitTyped > blueMax {
						blueMax = digitTyped
					}
				}
			}

			result += redMax * greenMax * blueMax
		}

		return
	}

	input := utils.ReadInput("day-02/input")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
