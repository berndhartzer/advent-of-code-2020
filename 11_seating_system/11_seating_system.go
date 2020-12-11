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
