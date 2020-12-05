package binary_boarding

import (
	"sort"
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

func BinaryBoardingPartTwo(passes []string) int64 {
	taken := make([]int64, len(passes))

	for i, pass := range passes {
		bin := strings.ReplaceAll(pass, "F", "0")
		bin = strings.ReplaceAll(bin, "B", "1")
		bin = strings.ReplaceAll(bin, "L", "0")
		bin = strings.ReplaceAll(bin, "R", "1")

		num, _ := strconv.ParseInt(bin, 2, 64)
		taken[i] = num
	}

	sort.Slice(taken, func(a, b int) bool {
		return taken[a] < taken[b]
	})

	var empty int64 = 0
	for i := 1; i <= len(taken)-1; i++ {
		if taken[i] != taken[i-1]+1 {
			empty = taken[i] - 1
			break
		}
	}

	return empty
}
