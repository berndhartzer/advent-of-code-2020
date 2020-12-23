package crab_combat

import (
	"strconv"
)

func CrabCombatPartOne(cards []string) int {
	playerOne := []int{}
	playerTwo := []int{}

	player := 1
	for i := 1; i < len(cards); i++ {
		if cards[i] == "" {
			player = 2
			i++
			continue
		}
		n, _ := strconv.Atoi(cards[i])

		if player == 1 {
			playerOne = append(playerOne, n)
		} else {
			playerTwo = append(playerTwo, n)
		}
	}

	for len(playerOne) > 0 && len(playerTwo) > 0 {
		playerOneCard := playerOne[0]
		playerTwoCard := playerTwo[0]

		var winner int
		var prize int
		if playerOneCard > playerTwoCard {
			winner, playerOne = playerOne[0], playerOne[1:]
			prize, playerTwo = playerTwo[0], playerTwo[1:]

			playerOne = append(playerOne, winner, prize)
		} else {
			winner, playerTwo = playerTwo[0], playerTwo[1:]
			prize, playerOne = playerOne[0], playerOne[1:]

			playerTwo = append(playerTwo, winner, prize)
		}
	}

	if len(playerOne) > 0 {
		return calculateScore(playerOne)
	} else {
		return calculateScore(playerTwo)
	}
}

func calculateScore(cards []int) int {
	m := len(cards)
	total := 0

	for _, card := range cards {
		total += (card * m)
		m--
	}

	return total
}
