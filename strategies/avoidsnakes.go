package strategies

import "github.com/roca/battlesnake/types"

func AvoidSnakes(board types.Board, you types.Battlesnake, nextmove string) bool {
	futureHead := predict_future_position(you.Head, nextmove)
	for _, snake := range board.Snakes {
		if !avoidSnake(futureHead, snake) && you.Length <= snake.Length {
			return false
		}
	}
	return true
}

func avoidSnake(futureHead types.Coord, snake types.Battlesnake) bool {
	snakeTail := snake.Body[snake.Length-1]

	for _, coord := range snake.Body {
		// Chasing tails are valid moves
		if futureHead.X == snakeTail.X && futureHead.Y == snakeTail.Y {
			return true
		}
		if futureHead.X == coord.X && futureHead.Y == coord.Y {
			return false
		}
	}
	return true
}
