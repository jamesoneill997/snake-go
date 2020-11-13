package main

import (
	"fmt"
	"strings"
)

/*
snake
controls
food
border
movement
hit logic
score
*/

//Snake is the player struct
type Snake struct {
	length string
	alive  bool
	speed  float32
}

//DrawBox function will draw the gameboard
func DrawBox(w int, h int, snakePos [2]int, foodPos [2]int) {
	for i := 0; i <= h; i++ {
		if i == 0 || i == h {
			fmt.Println(strings.Repeat("-", w))
		} else {
			fmt.Println("|", strings.Repeat(" ", w-4), "|")
		}
	}
}

func main() {
	boardSize := [2]int{124, 80} //h,w - Ratio of 31:20
	snakePos := [2]int{0, 0}
	foodPos := [2]int{10, 10}

	DrawBox(boardSize[0], boardSize[1], snakePos, foodPos)

}
