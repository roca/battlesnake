package strategies

import (
	"github.com/roca/battlesnake/types"
)

func AvoidWalls(board types.Board, you types.Battlesnake, next_move string) bool {
	futureHead := predict_future_position(you.Head, next_move)

	if futureHead.X < 0 || futureHead.Y < 0 {
		return false
	}
	if futureHead.X >= board.Width {
		return false
	}
	if futureHead.Y >= board.Height {
		return false
	}

	return true
}

func predict_future_position(currentHead types.Coord, next_move string) types.Coord {
	futureHead := currentHead

	switch next_move {
	case "left":
		// moving left means decreasing X by 1
		futureHead.X -= 1
	case "right":
		// moving right means increasing X by 1
		futureHead.X += 1
	case "up":
		// moving up means increasing Y by 1
		futureHead.Y += 1
	case "down":
		// moving down means decreasing Y by 1
		futureHead.Y -= 1
	}

	return futureHead
}
