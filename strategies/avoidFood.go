package strategies

import "github.com/roca/battlesnake/types"

func AvoidFood(board types.Board, you types.Battlesnake, next_move string) bool {
	futureHead := predict_future_position(you.Head, next_move)
	for _, coord := range board.Food {
		if futureHead.X == coord.X && futureHead.Y == coord.Y && you.Length > 5 {
			return false
		}
	}
	return true
}
