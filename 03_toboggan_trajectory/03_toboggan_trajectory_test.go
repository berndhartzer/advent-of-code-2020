package toboggan_trajectory

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestTobogganTrajectoryPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []string{
			"..##.......",
			"#...#...#..",
			".#....#..#.",
			"..#.#...#.#",
			".#...##..#.",
			"..#.##.....",
			".#.#.#....#",
			".#........#",
			"#.##...#...",
			"#...##....#",
			".#..#...#.#",
		}

		expected := 7
		got := TobogganTrajectoryPartOne(example, 3, 1)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./03_input.txt")

		output := TobogganTrajectoryPartOne(input, 3, 1)
		t.Log(fmt.Sprintf("%d\n", output))
	})
}

func TestTobogganTrajectoryPartTwo(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []string{
			"..##.......",
			"#...#...#..",
			".#....#..#.",
			"..#.#...#.#",
			".#...##..#.",
			"..#.##.....",
			".#.#.#....#",
			".#........#",
			"#.##...#...",
			"#...##....#",
			".#..#...#.#",
		}
		movements := [][]int{
			{1, 1},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		}

		expected := 336
		got := TobogganTrajectoryPartTwo(example, movements)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./03_input.txt")
		movements := [][]int{
			{1, 1},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		}

		output := TobogganTrajectoryPartTwo(input, movements)
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
