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
