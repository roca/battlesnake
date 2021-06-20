package test

import (
	"testing"

	"github.com/roca/battlesnake/strategies"
	"github.com/roca/battlesnake/types"
	"github.com/stretchr/testify/assert"
)

func TestAvoidWalls(t *testing.T) {
	snake := types.Battlesnake{
		Head: types.Coord{X: 0, Y: 0},
	}
	move := "up"
	board := types.Board{
		Width:  1,
		Height: 2,
	}
	safeMove := true
	actual := strategies.AvoidWalls(board, snake, move)
	assert.Equalf(t, safeMove, actual, "TestAvoidWalls: safeMove%f != actual %f", safeMove, actual)

	move = "down"
	safeMove = false
	actual = strategies.AvoidWalls(board, snake, move)
	assert.Equalf(t, safeMove, actual, "TestAvoidWalls: safeMove%f != actual %f", safeMove, actual)

	move = "left"
	safeMove = false
	actual = strategies.AvoidWalls(board, snake, move)
	assert.Equalf(t, safeMove, actual, "TestAvoidWalls: safeMove%f != actual %f", safeMove, actual)

	move = "right"
	safeMove = false
	actual = strategies.AvoidWalls(board, snake, move)
	assert.Equalf(t, safeMove, actual, "TestAvoidWalls: safeMove%f != actual %f", safeMove, actual)

}
