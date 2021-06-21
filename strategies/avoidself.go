package strategies

import (
	"github.com/roca/battlesnake/types"
)

func AvoidSelf(you types.Battlesnake, nextmove string) bool {
	futureHead := predict_future_position(you.Head, nextmove)
	return avoidSnake(futureHead, you)
}
