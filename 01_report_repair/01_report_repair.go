package report_repair

const target = 2020

func ReportRepairPartOne(entries []int) int {
	lookup := make(map[int]int)

	for i, v := range entries {
		diff := target - v

		j, ok := lookup[diff]
		if ok {
			return entries[j] * entries[i]
		}

		lookup[v] = i
	}

	return 0
}
