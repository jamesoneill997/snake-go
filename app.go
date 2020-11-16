package main

import (
	"log"

	. "github.com/gbin/goncurses"
)

func KeyInSlice(a Key, list [4]Key) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Insert(a [][]int, index int, value []int) [][]int {

	//if index is last element
	if len(a) == index {
		return append(a, value)
	}

	//else, split array at index and add value in the middle
	a = append(a[:index+1], a[index:]...)
	a[index] = value

	return a
}

func main() {

	stdscr, err := Init()
	if err != nil {
		log.Fatal(err)
	}
	defer End()
	Raw(true)
	Echo(false)
	Cursor(0)
	stdscr.Clear()
	stdscr.Keypad(true)

	my, mx := stdscr.MaxYX()

	//head := [2]int{3, 3}
	snake := [][]int{{my / 2, mx/2 + 1}, {my / 2, mx / 2}, {my / 2, mx/2 - 1}}
	var direction Key

	win, _ := NewWindow(40, 120, 10, 35)
	win.Keypad(true)

	for y := range snake {
		stdscr.MovePrint(snake[y][0], snake[y][1], "o")
	}

	for {
		keys := [4]Key{KEY_UP, KEY_DOWN, KEY_LEFT, KEY_RIGHT}
		key := stdscr.GetChar()

		if KeyInSlice(key, keys) {
			direction = key
		} else {
			direction = direction
		}
		head := snake[0]
		newHead := head

		if direction == keys[3] {
			newHead[0], newHead[1] = head[0], head[1]+1

		} else if direction == keys[2] {
			newHead[0], newHead[1] = head[0], head[1]-1
		} else if direction == keys[1] {
			newHead[0], newHead[1] = head[0]+1, head[1]
		} else if direction == keys[0] {
			newHead[0], newHead[1] = head[0]-1, head[1]
		}

		snake = Insert(snake, 0, newHead)
		stdscr.MovePrint(newHead[0], newHead[1], "o")
		stdscr.MovePrint(snake[len(snake)-1][0], snake[len(snake)-1][1], " ")
		tail := len(snake)
		snake = append(snake[:tail])
		stdscr.Refresh()
	}

	stdscr.Print("Use arrow keys to go up and down, Press enter to select")
	stdscr.Refresh()
}
