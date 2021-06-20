package test

import (
	"testing"

	"github.com/roca/battlesnake/strategies"
	"github.com/roca/battlesnake/types"
	"github.com/stretchr/testify/assert"
)

func TestAvoidFood(t *testing.T) {
	snake := types.Battlesnake{
		Head:   types.Coord{X: 0, Y: 0},
		Health: 99,
		Length: 9,
	}
	move := "up"
	board := types.Board{
		Food: []types.Coord{
			{X: 1, Y: 0},
			{X: 0, Y: 1},
			{X: 1, Y: 1},
		},
	}
	safeMove := false
	actual := strategies.AvoidFood(board, snake, move)
	assert.Equalf(t, safeMove, actual, "TestAvoidFood: safeMove%f != actual %f", safeMove, actual)

	snake.Health = 19 // Health below min of 20
	safeMove = true
	actual = strategies.AvoidFood(board, snake, move)
	assert.Equalf(t, safeMove, actual, "TestAvoidFood: safeMove %f != actual %f", safeMove, actual)

}
