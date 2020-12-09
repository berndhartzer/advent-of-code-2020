package encoding_error

func EncodingErrorPartOne(nums []int, preamble int) int {
	curr := 0

	for i := preamble; i < len(nums); i++ {
		curr = nums[i]
		check := nums[i-preamble : i]
		valid := false

		lookup := make(map[int]int)
		for j, v := range check {
			diff := curr - v

			_, ok := lookup[diff]
			if ok {
				valid = true
				break
			}

			lookup[v] = j
		}

		if !valid {
			break
		}
	}

	return curr
}
