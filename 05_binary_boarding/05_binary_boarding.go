package binary_boarding

import (
	"strconv"
	"strings"
)

func BinaryBoardingPartOne(passes []string) int64 {
	var highest int64 = 0

	for _, pass := range passes {
		bin := strings.ReplaceAll(pass, "F", "0")
		bin = strings.ReplaceAll(bin, "B", "1")
		bin = strings.ReplaceAll(bin, "L", "0")
		bin = strings.ReplaceAll(bin, "R", "1")

		num, _ := strconv.ParseInt(bin, 2, 64)
		if num > highest {
			highest = num
		}
	}

	return highest
}
