package strategies

import (
	"github.com/roca/battlesnake/types"
)

func AvoidWalls(board types.Board, currentHead types.Coord, next_move string) bool {
	futureHead := predict_future_position(currentHead, next_move)

	if futureHead.X < 0 || futureHead.Y < 0 {
		return false
	}
	if futureHead.X > board.Width {
		return false
	}
	if futureHead.Y > board.Height {
		return false
	}

	return true
}

func predict_future_position(currentHead types.Coord, next_move string) types.Coord {
	futureHead := currentHead

	switch next_move {
	case "left":
		// moving left means increasing x by 1
		futureHead.X = currentHead.X - 1
	case "right":
		// moving right means decreasing x by 1
		futureHead.X = currentHead.X + 1
	case "up":
		// moving up means increasing Y by 1
		futureHead.Y = currentHead.Y + 1
	case "down":
		// moving down means decreasing Y by 1
		futureHead.Y = currentHead.Y - 1
	}

	return futureHead
}
