package handy_haversacks

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestHandyHaversacksPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []string{
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
			"bright white bags contain 1 shiny gold bag.",
			"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
			"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
			"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
			"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
			"faded blue bags contain no other bags.",
			"dotted black bags contain no other bags.",
		}
		target := "shiny gold"

		expected := 4
		got := HandyHaversacksPartOne(example, target)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./07_input.txt")
		target := "shiny gold"

		output := HandyHaversacksPartOne(input, target)
		t.Log(fmt.Sprintf("%d\n", output))
	})
}

func TestHandyHaversacksPartTwo(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []string{
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
			"bright white bags contain 1 shiny gold bag.",
			"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
			"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
			"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
			"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
			"faded blue bags contain no other bags.",
			"dotted black bags contain no other bags.",
		}
		target := "shiny gold"

		expected := 32
		got := HandyHaversacksPartTwo(example, target)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("example input two", func(t *testing.T) {
		example := []string{
			"shiny gold bags contain 2 dark red bags.",
			"dark red bags contain 2 dark orange bags.",
			"dark orange bags contain 2 dark yellow bags.",
			"dark yellow bags contain 2 dark green bags.",
			"dark green bags contain 2 dark blue bags.",
			"dark blue bags contain 2 dark violet bags.",
			"dark violet bags contain no other bags.",
		}
		target := "shiny gold"

		expected := 126
		got := HandyHaversacksPartTwo(example, target)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./07_input.txt")
		target := "shiny gold"

		output := HandyHaversacksPartTwo(input, target)
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
