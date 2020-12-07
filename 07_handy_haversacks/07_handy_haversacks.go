package handy_haversacks

import (
	"regexp"
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
			containsMatches := regexp.MustCompile("(\\d+) (.+) bag").FindStringSubmatch(part)
			containee := containsMatches[2]

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
