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
