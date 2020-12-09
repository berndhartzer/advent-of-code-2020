package encoding_error

func EncodingErrorPartOne(nums []int, preamble int) (int, int) {
	curr := 0
	i := preamble

	for ; i < len(nums); i++ {
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

	return curr, i
}

func EncodingErrorPartTwo(nums []int, preamble int) int {
	invalid, limit := EncodingErrorPartOne(nums, preamble)

	check := nums[0:limit]
	contiguous := []int{}

outer:
	for i := range check {
		sub := check[i:]

		acc := 0
		tmp := []int{}
		for _, n := range sub {
			tmp = append(tmp, n)
			acc += n

			if acc == invalid {
				contiguous = tmp
				break outer
			}
		}
	}

	lowest := 0
	highest := 0
	for _, n := range contiguous {
		if n > highest {
			highest = n
		}

		if n < lowest || lowest == 0 {
			lowest = n
		}
	}

	return highest + lowest
}
