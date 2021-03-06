package adapter_array

import (
	"sort"
)

func AdapterArrayPartOne(adapters []int) int {
	differences := make(map[int]int)

	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	joltage := 0
	for _, n := range adapters {
		diff := n - joltage
		differences[diff]++

		joltage = n
	}

	return differences[1] * differences[3]
}

func AdapterArrayPartTwo(adapters []int) int {
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	variations := make(map[int]int)
	variations[0] = 1
	for _, n := range adapters {
		variations[n] = variations[n-1] + variations[n-2] + variations[n-3]
	}

	return variations[adapters[len(adapters)-1]]
}
