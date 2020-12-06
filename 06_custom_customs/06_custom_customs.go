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
