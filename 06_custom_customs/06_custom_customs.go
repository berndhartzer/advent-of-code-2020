package custom_customs

import (
	"strings"
)

func CustomCustomsPartOne(forms []string) int {
	counts := 0
	answers := make(map[string]bool)

	for i, form := range forms {
		if form == "" {
			continue
		}

		split := strings.Split(form, "")
		for _, l := range split {
			answers[l] = true
		}

		if i >= len(forms)-1 || forms[i+1] == "" {
			for k, v := range answers {
				if v {
					counts++
				}

				answers[k] = false
			}
		}
	}

	return counts
}

func CustomCustomsPartTwo(forms []string) int {
	counts := 0

	answers := make(map[string]int)
	group := 0

	for i, form := range forms {
		if form == "" {
			continue
		}

		group++

		split := strings.Split(form, "")
		for _, l := range split {
			answers[l]++
		}

		if i >= len(forms)-1 || forms[i+1] == "" {
			for k, v := range answers {
				if v == group {
					counts++
				}

				answers[k] = 0
			}

			group = 0
		}
	}

	return counts
}
