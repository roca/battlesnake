package strategies

import "github.com/roca/battlesnake/types"

func AvoidSnakes(board types.Board, you types.Battlesnake, next_move string) bool {
	futureHead := predict_future_position(you.Head, next_move)
	// If futureHead == self current tail
	// this should be safe
	selfCurrentTail := you.Body[you.Length-1]

	for _, snake := range board.Snakes {

		if futureHead.X == selfCurrentTail.X && futureHead.Y == selfCurrentTail.Y {
			return true
		}

		if !avoidSnake(futureHead, snake) {
			return false
		}
	}

	return true
}

func avoidSnake(futureHead types.Coord, snake types.Battlesnake) bool {
	for _, coord := range snake.Body {
		if futureHead.X == coord.X && futureHead.Y == coord.Y {
			return false
		}
	}
	return true
}
