package seating_system

const empty byte = 76
const occupied byte = 35
const floor byte = 46

func SeatingSystemPartOne(seats []string) int {
	state := make(map[int][]byte)
	for y := 0; y < len(seats); y++ {
		for x := 0; x < len(seats[y]); x++ {
			state[y] = append(state[y], seats[y][x])
		}
	}

	for {
		changed := false

		newState := make(map[int][]byte)
		for y, v := range state {
			newState[y] = make([]byte, len(v))
			for x, pos := range v {
				newState[y][x] = pos

				if pos == floor {
					continue
				}

				adjacentOccupied := 0

				rows := []int{y - 1, y + 1}
				for _, r := range rows {
					row, ok := state[r]
					if ok {
						if x-1 >= 0 {
							if row[x-1] == occupied {
								adjacentOccupied++
							}
						}
						if row[x] == occupied {
							adjacentOccupied++
						}
						if x+1 <= len(row)-1 {
							if row[x+1] == occupied {
								adjacentOccupied++
							}
						}
					}
				}

				if x-1 >= 0 {
					if state[y][x-1] == occupied {
						adjacentOccupied++
					}
				}

				if x+1 <= len(state[y])-1 {
					if state[y][x+1] == occupied {
						adjacentOccupied++
					}
				}

				switch pos {
				case empty:
					if adjacentOccupied == 0 {
						newState[y][x] = occupied

						if pos != occupied {
							changed = true
						}
					}

				case occupied:
					if adjacentOccupied >= 4 {
						newState[y][x] = empty

						if pos != empty {
							changed = true
						}
					}
				}
			}
		}

		if !changed {
			break
		}

		state = newState
	}

	count := 0
	for _, v := range state {
		for _, p := range v {
			if p == occupied {
				count++
			}
		}
	}

	return count
}

type Grid struct {
	state map[int][]byte
}

func (g Grid) NumVisibleInColumn(x, y, direction int) int {
	count := 0

	check := y + direction
	stopSearching := false

	for {
		row, ok := g.state[check]
		if ok {
			switch row[x] {
			case empty:
				stopSearching = true
			case occupied:
				stopSearching = true
				count++
			}
		} else {
			stopSearching = true
		}

		check += direction

		if stopSearching {
			break
		}
	}

	return count
}

func (g Grid) NumVisibleInRow(x, y, direction int) int {
	count := 0

	check := x + direction
	stopSearching := false

	for {
		if check >= 0 && check <= len(g.state[y])-1 {
			switch g.state[y][check] {
			case empty:
				stopSearching = true
			case occupied:
				stopSearching = true
				count++
			}
		} else {
			stopSearching = true
		}

		check += direction

		if stopSearching {
			break
		}
	}

	return count
}

func (g Grid) NumVisibleInDiagonal(x, y, dirX, dirY int) int {
	count := 0

	checkY := y + dirY
	checkX := x + dirX
	stopSearching := false

	for {
		_, ok := g.state[checkY]
		if ok {
			if checkX >= 0 && checkX <= len(g.state[checkY])-1 {
				switch g.state[checkY][checkX] {
				case empty:
					stopSearching = true
				case occupied:
					stopSearching = true
					count++
				}
			} else {
				stopSearching = true
			}
		} else {
			stopSearching = true
		}

		checkY += dirY
		checkX += dirX

		if stopSearching {
			break
		}
	}

	return count
}

func (g Grid) NumVisibleOccupied(x, y int) int {
	count := 0

	count += g.NumVisibleInColumn(x, y, -1)
	count += g.NumVisibleInDiagonal(x, y, 1, -1)
	count += g.NumVisibleInRow(x, y, 1)
	count += g.NumVisibleInDiagonal(x, y, 1, 1)
	count += g.NumVisibleInColumn(x, y, 1)
	count += g.NumVisibleInDiagonal(x, y, -1, 1)
	count += g.NumVisibleInRow(x, y, -1)
	count += g.NumVisibleInDiagonal(x, y, -1, -1)

	return count
}

func SeatingSystemPartTwo(seats []string) int {
	grid := Grid{
		state: make(map[int][]byte),
	}

	for y := 0; y < len(seats); y++ {
		for x := 0; x < len(seats[y]); x++ {
			grid.state[y] = append(grid.state[y], seats[y][x])
		}
	}

	for {
		changed := false

		newState := make(map[int][]byte)
		for y, v := range grid.state {
			newState[y] = make([]byte, len(v))
			for x, pos := range v {
				newState[y][x] = pos

				if pos == floor {
					continue
				}

				adjacentOccupied := grid.NumVisibleOccupied(x, y)

				switch pos {
				case empty:
					if adjacentOccupied == 0 {
						newState[y][x] = occupied

						if pos != occupied {
							changed = true
						}
					}

				case occupied:
					if adjacentOccupied >= 5 {
						newState[y][x] = empty

						if pos != empty {
							changed = true
						}
					}
				}
			}
		}

		if !changed {
			break
		}

		grid.state = newState
	}

	count := 0
	for _, v := range grid.state {
		for _, p := range v {
			if p == occupied {
				count++
			}
		}
	}

	return count
}
