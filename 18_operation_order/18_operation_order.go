package operation_order

import (
	"regexp"
	"strconv"
)

const space rune = 32
const openBracket rune = 40
const closeBracket rune = 41
const multiply rune = 42
const plus rune = 43

func evaluate(equation string) int {
	bracketTotal := make(map[int]int)
	bracketOp := make(map[int]rune)
	brackets := 0

	token := 0

	captureToken := func(i int) {
		if token > 0 {
			n, _ := strconv.Atoi(equation[i-token : i])
			if bracketOp[brackets] == multiply {
				bracketTotal[brackets] *= n
			} else {
				bracketTotal[brackets] += n
			}

			token = 0
		}
	}

	for i, char := range equation {
		switch char {
		case space:
			captureToken(i)

		case openBracket:
			brackets++

		case closeBracket:
			captureToken(i)

			bracketFinal := bracketTotal[brackets]
			bracketTotal[brackets] = 0
			bracketOp[brackets] = ' '

			brackets--

			if bracketOp[brackets] == multiply {
				bracketTotal[brackets] *= bracketFinal
			} else {
				bracketTotal[brackets] += bracketFinal
			}

		case multiply:
			bracketOp[brackets] = multiply

		case plus:
			bracketOp[brackets] = plus

		default:
			token++
		}

		if i >= len(equation)-1 {
			captureToken(i + 1)
		}
	}

	return bracketTotal[0]
}

func OperationOrderPartOne(equations []string) int {
	total := 0

	for _, equation := range equations {
		total += evaluate(equation)
	}

	return total
}

func OperationOrderPartTwo(equations []string) int {
	total := 0

	rePlus := regexp.MustCompile(`\d+ \+ \d+`)
	reBrackets := regexp.MustCompile(`\([^\(\)]+\)`)

	for _, equation := range equations {
		precedence := true

		for precedence {
			precedence = false

			equation = rePlus.ReplaceAllStringFunc(equation, func(s string) string {
				precedence = true

				evaluated := evaluate(s)
				return strconv.Itoa(evaluated)
			})

			if !precedence {
				equation = reBrackets.ReplaceAllStringFunc(equation, func(s string) string {
					precedence = true

					evaluated := evaluate(s)
					return strconv.Itoa(evaluated)
				})
			}
		}

		total += evaluate(equation)
	}

	return total
}
