package passport_processing

import (
	"regexp"
	"strconv"
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

func PassportProcessingPartTwo(passports []string) int {
	numValid := 0

	splitter := func(r rune) bool {
		return r == ' ' || r == ':'
	}

	validECL := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	fields := map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
	}

	for i, line := range passports {
		if line == "" {
			continue
		}

		split := strings.FieldsFunc(line, splitter)

		for j := 0; j < len(split); j += 2 {
			part := split[j]
			strValue := split[j+1]
			valid := false

			switch part {
			case "byr":
				v, err := strconv.Atoi(strValue)
				if err == nil {
					if v >= 1920 && v <= 2002 {
						valid = true
					}
				}

			case "iyr":
				v, err := strconv.Atoi(strValue)
				if err == nil {
					if v >= 2010 && v <= 2020 {
						valid = true
					}
				}

			case "eyr":
				v, err := strconv.Atoi(strValue)
				if err == nil {
					if v >= 2020 && v <= 2030 {
						valid = true
					}
				}

			case "hgt":
				measurement := strValue[len(strValue)-2:]
				v, err := strconv.Atoi(strValue[:len(strValue)-2])
				if err == nil {
					if measurement == "cm" {
						if v >= 150 && v <= 193 {
							valid = true
						}
					} else if measurement == "in" {
						if v >= 59 && v <= 76 {
							valid = true
						}
					}
				}

			case "hcl":
				match, _ := regexp.MatchString("^#([a-f0-9]{6})$", strValue)
				if match {
					valid = true
				}

			case "ecl":
				for _, c := range validECL {
					if strValue == c {
						valid = true
					}
				}

			case "pid":
				match, _ := regexp.MatchString("^[0-9]{9}$", strValue)
				if match {
					valid = true
				}
			}

			fields[part] = valid
		}

		if i >= len(passports)-1 || passports[i+1] == "" {
			_, ok := fields["cid"]
			if ok {
				delete(fields, "cid")
			}

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
		}
	}

	return numValid
}
