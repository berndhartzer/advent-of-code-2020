package toboggan_trajectory

const tree byte = 35

func TobogganTrajectoryPartOne(mapLines []string, right, down int) int {
	trees := 0
	x, y := 0, 0

	for y <= len(mapLines)-1 {
		if x > len(mapLines[y])-1 {
			x = x - len(mapLines[y])
		}

		pos := mapLines[y][x]
		if pos == tree {
			trees++
		}

		x += right
		y += down
	}

	return trees
}

func TobogganTrajectoryPartTwo(mapLines []string, movements [][]int) int {
	trees := []int{}

	for _, movement := range movements {
		trees = append(trees, TobogganTrajectoryPartOne(mapLines, movement[0], movement[1]))
	}

	total := trees[0]

	for i := 1; i <= len(trees)-1; i++ {
		total *= trees[i]
	}

	return total
}
