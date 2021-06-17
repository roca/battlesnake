package strategies

import (
	"fmt"
	"math/rand"

	"github.com/robpike/filter"
	"github.com/roca/battlesnake/types"
)

func AvoidSelfWallsAndFood(game types.GameRequest) string {
	//Choose a random direction to move in
	possibleMoves := []string{"up", "down", "left", "right"}
	move := possibleMoves[rand.Intn(len(possibleMoves))]
	for i := 0; i < 4; i++ {

		safe := AvoidSelf(game.You, move) &&
			AvoidWalls(game.Board, game.You, move) &&
			AvoidFood(game.Board, game.You, move)
		fmt.Println(move, "Attempt", i, "AvoidSnakesAndWalls:", safe)
		if safe {
			break
		}
		// Drop unsafe move from possible moves
		possibleMoves = filter.Drop(possibleMoves, func(m string) bool {
			return m == move
		}).([]string)
		if len(possibleMoves) == 0 {
			break
		}
		move = possibleMoves[rand.Intn(len(possibleMoves))]
	}
	return move
}

func AvoidSnakesAndWalls(game types.GameRequest) string {
	//Choose a random direction to move in
	possibleMoves := []string{"up", "down", "left", "right"}
	move := possibleMoves[rand.Intn(len(possibleMoves))]
	for i := 0; i < 4; i++ {

		safe := AvoidSnakes(game.Board, game.You, move) && AvoidWalls(game.Board, game.You, move)
		fmt.Println(move, "Attempt", i, "AvoidSnakesAndWalls:", safe)
		if safe {
			break
		}
		// Drop unsafe move from possible moves
		possibleMoves = filter.Drop(possibleMoves, func(m string) bool {
			return m == move
		}).([]string)
		if len(possibleMoves) == 0 {
			break
		}
		move = possibleMoves[rand.Intn(len(possibleMoves))]
	}
	return move
}

func AvoidSnakesWallsAndFood(game types.GameRequest) string {
	//Choose a random direction to move in
	possibleMoves := []string{"up", "down", "left", "right"}
	move := possibleMoves[rand.Intn(len(possibleMoves))]
	for i := 0; i < 4; i++ {

		safe := AvoidSnakes(game.Board, game.You, move) &&
			AvoidWalls(game.Board, game.You, move) &&
			AvoidFood(game.Board, game.You, move)
		fmt.Println(move, "Attempt", i, "AvoidSnakesWallsAndFood:", safe)
		if safe {
			break
		}
		// Drop unsafe move from possible moves
		possibleMoves = filter.Drop(possibleMoves, func(m string) bool {
			return m == move
		}).([]string)
		if len(possibleMoves) == 0 {
			break
		}
		move = possibleMoves[rand.Intn(len(possibleMoves))]
	}
	return move

}
