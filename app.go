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
*/

func DrawBox() {
	for i := 0; i <= 20; i++ {
		if i == 0 || i == 20 {
			fmt.Println(strings.Repeat("-", 31))
		} else {
			fmt.Println("|", strings.Repeat(" ", 27), "|")
		}
	}
}

func main() {
	DrawBox()
}
