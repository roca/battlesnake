package test

import (
	"testing"

	"github.com/roca/battlesnake/strategies"
	"github.com/roca/battlesnake/types"
	"github.com/stretchr/testify/assert"
)

func TestAvoidSnakes(t *testing.T) {

	you := types.Battlesnake{
		Head: types.Coord{X: 3, Y: 1},
		Body: []types.Coord{
			{X: 3, Y: 1},
			{X: 4, Y: 1},
			{X: 4, Y: 2},
			{X: 4, Y: 3},
			{X: 4, Y: 4},
		},
		Length: 4,
	}
	snake1 := types.Battlesnake{
		Head: types.Coord{X: 1, Y: 1},
		Body: []types.Coord{
			{X: 3, Y: 2},
			{X: 3, Y: 3},
			{X: 3, Y: 4},
			{X: 3, Y: 5},
			{X: 3, Y: 6},
			{X: 3, Y: 7},
		},
		Length: 5,
	}

	board := types.Board{
		Snakes: []types.Battlesnake{you, snake1},
	}
	move := "up"
	safeMove := false
	actual := strategies.AvoidSnakes(board, you, move)
	assert.Equalf(t, safeMove, actual, "TestAvoidSnakes safeMove%f != actual %f", safeMove, actual)

	move = "left"
	safeMove = true
	actual = strategies.AvoidSnakes(board, you, move)
	assert.Equalf(t, safeMove, actual, "TestAvoidSnakes: safeMove%f != actual %f", safeMove, actual)


}
