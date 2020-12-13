package shuttle_search

import (
	"strconv"
)

func ShuttleSearchPartOne(earliest int, buses []string) int {
	lowest := 0
	bus := 0

	for _, b := range buses {
		if b == "x" {
			continue
		}

		busNum, _ := strconv.Atoi(b)

		multiples := earliest / busNum

		next := busNum * (multiples + 1)

		diff := next - earliest
		if diff < lowest || lowest == 0 {
			lowest = diff
			bus = busNum
		}
	}

	return lowest * bus
}
