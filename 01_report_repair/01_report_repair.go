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

func ReportRepairPartTwo(entries []int) int {
	for i, ii := range entries {
		for j, jj := range entries[i+1:] {
			for _, kk := range entries[j+1:] {
				if ii+jj+kk == target {
					return ii * jj * kk
				}
			}
		}
	}

	return 0
}
