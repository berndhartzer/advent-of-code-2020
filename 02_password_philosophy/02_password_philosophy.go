package password_philosophy

import (
	"strconv"
	"strings"
)

func PasswordPhilosophyPartOne(entries []string) int {
	valid := 0

	for _, v := range entries {
		split := strings.Split(v, " ")

		params := strings.Split(split[0], "-")
		lower, _ := strconv.Atoi(params[0])
		upper, _ := strconv.Atoi(params[1])

		target := strings.TrimSuffix(split[1], ":")
		password := split[2]

		num := strings.Count(password, target)
		if num >= lower && num <= upper {
			valid++
		}
	}

	return valid
}

func PasswordPhilosophyPartTwo(entries []string) int {
	valid := 0

	for _, v := range entries {
		split := strings.Split(v, " ")

		params := strings.Split(split[0], "-")
		first, _ := strconv.Atoi(params[0])
		second, _ := strconv.Atoi(params[1])

		target := strings.TrimSuffix(split[1], ":")
		password := split[2]

		matches := 0
		if string(password[first-1]) == target {
			matches++
		}
		if string(password[second-1]) == target {
			matches++
		}

		if matches == 1 {
			valid++
		}
	}

	return valid
}
