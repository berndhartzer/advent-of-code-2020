package ticket_translation

import (
	"regexp"
	"strconv"
	"strings"
)

func TicketTranslationPartOne(notes []string) int {
	i := 0
	tickets := []int{}
	rules := make(map[int]bool)

	// rules
	for ; i < len(notes); i++ {
		if notes[i] == "" {
			break
		}

		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(notes[i], -1)

		for j := 0; j < len(nums); j += 2 {
			lower, _ := strconv.Atoi(nums[j])
			upper, _ := strconv.Atoi(nums[j+1])

			for k := lower; k <= upper; k++ {
				rules[k] = true
			}
		}
	}
	i += 2

	// your ticket
	for ; i < len(notes); i++ {
		if notes[i] == "" {
			break
		}
	}
	i += 2

	// nearby tickets
	for ; i < len(notes); i++ {
		split := strings.Split(notes[i], ",")
		for _, v := range split {
			n, _ := strconv.Atoi(v)
			tickets = append(tickets, n)
		}
	}

	invalid := 0
	for _, ticket := range tickets {
		_, ok := rules[ticket]
		if !ok {
			invalid += ticket
		}
	}

	return invalid
}
