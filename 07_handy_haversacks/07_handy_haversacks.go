package handy_haversacks

import (
	"regexp"
	"strconv"
	"strings"
)

func HandyHaversacksPartOne(rules []string, target string) int {
	all := map[string][]string{}

	for _, rule := range rules {
		split := strings.Split(rule, " bags contain ")
		if strings.Contains(split[1], "no other bags") {
			continue
		}

		name := split[0]

		containsSplit := strings.Split(split[1], ",")
		for _, part := range containsSplit {
			containsBags := regexp.MustCompile("(\\d+) (.+) bag").FindStringSubmatch(part)
			containee := containsBags[2]

			_, ok := all[containee]
			if !ok {
				all[containee] = []string{}
			}

			all[containee] = append(all[containee], name)
		}
	}

	containers := all[target]
	for i := 0; i <= len(containers)-1; i++ {
		v, ok := all[containers[i]]
		if ok {
			containers = append(containers, v...)
		}
	}

	unique := make(map[string]bool)
	for _, c := range containers {
		unique[c] = true
	}

	return len(unique)
}

type Bag struct {
	name     string
	contains map[string]int
}

func HandyHaversacksPartTwo(rules []string, target string) int {
	all := make(map[string]*Bag)

	for _, rule := range rules {
		split := strings.Split(rule, " bags contain ")
		if strings.Contains(split[1], "no other bags") {
			continue
		}

		name := split[0]

		bag := &Bag{
			name:     name,
			contains: make(map[string]int),
		}

		containsSplit := strings.Split(split[1], ",")
		for _, part := range containsSplit {
			containsBags := regexp.MustCompile("(\\d+) (.+) bag").FindStringSubmatch(part)
			num, _ := strconv.Atoi(containsBags[1])
			containee := containsBags[2]

			bag.contains[containee] = num
		}

		all[name] = bag
	}

	bagsToCount := map[string]int{
		target: 1,
	}

	count := 0
	loop := true
	for loop {
		loop = false

		for k, v := range bagsToCount {
			if v > 0 {
				loop = true

				bag, ok := all[k]
				if ok {
					for kk, vv := range bag.contains {
						bagsToCount[kk] += vv * v
					}
				}

				count += v

				bagsToCount[k] = 0
			}
		}
	}

	return count - 1
}
