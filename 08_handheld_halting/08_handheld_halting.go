package handheld_halting

import (
	"strconv"
	"strings"
)

func HandheldHaltingPartOne(instructions []string) int {
	seen := make(map[int]bool)
	p := 0
	acc := 0

	for {
		_, ok := seen[p]
		if ok {
			break
		}
		seen[p] = true

		split := strings.Split(instructions[p], " ")
		num, _ := strconv.Atoi(split[1])

		switch split[0] {
		case "acc":
			acc += num
			p++

		case "jmp":
			p += num

		case "nop":
			p++
		}
	}

	return acc
}

func HandheldHaltingPartTwo(instructions []string) int {
	final := 0

outer:
	for i, instruction := range instructions {
		original, updated := "", ""

		if strings.Contains(instruction, "jmp") {
			original = "jmp"
			updated = "nop"
		} else if strings.Contains(instruction, "nop") {
			original = "nop"
			updated = "jmp"
		}

		if original != "" {
			tmp := make([]string, len(instructions))
			copy(tmp, instructions)

			tmp[i] = strings.Replace(tmp[i], original, updated, 1)

			seen := make(map[int]bool)
			p := 0
			acc := 0

			for {
				_, ok := seen[p]
				if ok {
					break
				}
				seen[p] = true

				split := strings.Split(tmp[p], " ")
				num, _ := strconv.Atoi(split[1])

				switch split[0] {
				case "acc":
					acc += num
					p++

				case "jmp":
					p += num

				case "nop":
					p++
				}

				if p >= len(tmp) {
					final = acc
					break outer
				}
			}
		}
	}

	return final
}
