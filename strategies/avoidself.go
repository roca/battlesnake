package strategies

import (
  "fmt"
	"github.com/roca/battlesnake/types"
)

func AvoidSelf(currentHead types.Coord, currentBody []types.Coord, next_move string) bool {
 futureHead := predict_future_position(currentHead, next_move)

  for _,coord := range currentBody{
    
    if futureHead.X == coord.X || futureHead.Y == coord.Y {
      fmt.Println("Ran into body",futureHead,coord)
      return false
    }
  }

  return true
}