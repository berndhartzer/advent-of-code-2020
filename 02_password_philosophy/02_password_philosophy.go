package password_philosophy

import (
	"fmt"
	"strings"
)

func PasswordPhilosophyPartOne(entries []string) int {
	valid := 0

	for _, v := range entries {
		// fmt.Println(v)

		split := strings.Split(v, " ")
		// fmt.Println(split)

		upperAndLower := split[0]

		target := strings.TrimSuffix(split[1], ":")
		password := split[2]
		fmt.Println(upperAndLower, target, password)

	}

	return valid
}
