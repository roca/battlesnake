package strategies

import (
	"github.com/roca/battlesnake/types"
)

func AvoidFood(board types.Board, you types.Battlesnake, nextMove string,options ...map[string]interface{}) bool {
	futureHead := predict_future_position(you.Head, nextMove)
	var minHealth int32 = 20
	var minLength int32 = 0
	if len(options) > 0 {
		minHealth = options[0]["minHealth"].(int32)
		minLength = options[0]["minLength"].(int32)
	}
	for _, coord := range board.Food {
		//fmt.Println(futureHead)
		if futureHead.X == coord.X &&
			futureHead.Y == coord.Y &&
			minHealth <= you.Health &&
			minLength <= you.Length  {
			return false
		}
	}
	return true
}
