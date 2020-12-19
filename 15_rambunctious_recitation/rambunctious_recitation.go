package rambunctious_recitation

func RambunctiousRecitation(nums []int, limit int) int {
	spoken := make(map[int][]int)
	last := 0

	for i, n := range nums {
		_, ok := spoken[n]
		if !ok {
			spoken[n] = []int{}
		}

		spoken[n] = append(spoken[n], i+1)
		last = n
	}

	for i := len(nums) + 1; i <= limit; i++ {
		v, ok := spoken[last]
		if !ok {
			spoken[last] = []int{}
		}

		newNum := 0
		if len(v) < 2 {
			newNum = nums[0]
			spoken[newNum] = append(spoken[newNum], i)
		} else {
			newNum = v[len(v)-1] - v[len(v)-2]

			_, oj := spoken[newNum]
			if !oj {
				spoken[newNum] = []int{}
			}
			spoken[newNum] = append(spoken[newNum], i)
		}

		last = newNum
	}

	return last
}
