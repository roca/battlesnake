package test

import (
	"testing"

	"github.com/roca/battlesnake/strategies"
	"github.com/roca/battlesnake/types"
	"github.com/stretchr/testify/assert"
)

func TestAvoidSelf(t *testing.T) {
	you := types.Battlesnake{
		Head: types.Coord{X: 1, Y: 1},
		Body: []types.Coord{
			{X: 1, Y: 1},
			{X: 2, Y: 1},
			{X: 2, Y: 2},
		},
		Length: 2,
	}
	/*
	        |2,2
	      ^ |
	   1,1--2,1
	*/
	move := "up"
	safeMove := true
	actual := strategies.AvoidSelf(you, move)
	assert.Equalf(t, safeMove, actual, "AvoidSelf: safeMove%f != actual %f", safeMove, actual)

	/*
	   1,2--|2,2
	      ^ |
	   1,1--2,1
	*/
	you.Body = append(you.Body, types.Coord{X:1,Y:2})
	safeMove = false
	actual = strategies.AvoidSelf(you, move)
	assert.Equalf(t, safeMove, actual, "AvoidSelf: safeMove%f != actual %f", safeMove, actual)

	/*
	   1,2--|2,2
	      < |
	   1,1--2,1
	*/
	move = "left"
	you.Body = append(you.Body, types.Coord{X:1,Y:2})
	safeMove = true
	actual = strategies.AvoidSelf(you, move)
	assert.Equalf(t, safeMove, actual, "AvoidSelf: safeMove%f != actual %f", safeMove, actual)
}
