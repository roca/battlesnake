package strategies

import (
	"github.com/roca/battlesnake/types"
)

func AvoidSelf(you types.Battlesnake, next_move string) bool {
	futureHead := predict_future_position(you.Head, next_move)
	return avoidSnake(futureHead, you)
}


