package passport_processing

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestPassportProcessingPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []string{
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
			"byr:1937 iyr:2017 cid:147 hgt:183cm",
			"",
			"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
			"hcl:#cfa07d byr:1929",
			"",
			"hcl:#ae17e1 iyr:2013",
			"eyr:2024",
			"ecl:brn pid:760753108 byr:1931",
			"hgt:179cm",
			"",
			"hcl:#cfa07d eyr:2025 pid:166559648",
			"iyr:2011 ecl:brn hgt:59in",
		}

		expected := 2
		got := PassportProcessingPartOne(example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./04_input.txt")

		output := PassportProcessingPartOne(input)
		t.Log(fmt.Sprintf("%d\n", output))
	})
}

func readFileAsStringSlice(t *testing.T, name string) []string {
	t.Helper()

	file, _ := os.Open(name)
	defer file.Close()

	var contents []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents
}
