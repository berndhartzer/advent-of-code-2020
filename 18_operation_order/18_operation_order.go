package operation_order

import (
	"strconv"
)

const space rune = 32
const openBracket rune = 40
const closeBracket rune = 41
const multiply rune = 42
const plus rune = 43

func OperationOrderPartOne(equations []string) int {
	total := 0

	for _, equation := range equations {
		bracketTotal := make(map[int]int)
		bracketOp := make(map[int]rune)
		brackets := 0

		for _, char := range equation {
			switch char {
			case space:
				continue

			case openBracket:
				brackets++

			case closeBracket:
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
				n, _ := strconv.Atoi(string(char))
				if bracketOp[brackets] == multiply {
					bracketTotal[brackets] *= n
				} else {
					bracketTotal[brackets] += n
				}
			}
		}

		total += bracketTotal[0]
	}

	return total
}
