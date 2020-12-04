package passport_processing

import (
	"strings"
)

func PassportProcessingPartOne(passports []string) int {
	numValid := 0

	fields := map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
	}

	for _, line := range passports {
		if line == "" {
			valid := true
			for _, v := range fields {
				if !v {
					valid = false
				}
			}

			if valid {
				numValid++
			}

			for field := range fields {
				fields[field] = false
			}

			continue
		}

		for field := range fields {
			if strings.Contains(line, field) {
				fields[field] = true
			}
		}
	}

	return numValid
}
