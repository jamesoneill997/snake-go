package main

import (
	"log"

	. "github.com/gbin/goncurses"
)

func main() {
	head := [2]int{3, 3}
	snake := []string{"0"}

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

	win, _ := NewWindow(40, 120, 10, 35)
	win.Keypad(true)

	stdscr.Print("Use arrow keys to go up and down, Press enter to select")
	stdscr.Refresh()
	printsnake(win, snake, head)

	for {
		ch := stdscr.GetChar()
		switch Key(ch) {
		case 'q':
			return
		case KEY_UP:
			head[0]--

		case KEY_DOWN:

			head[0]++

		case KEY_LEFT:

			head[1]--

		case KEY_RIGHT:
			head[1]++

		default:
			stdscr.ClearToEOL()
			stdscr.Refresh()

		}
		win.Erase()
		printsnake(win, snake, head)

	}
}

func printsnake(w *Window, snake []string, head [2]int) {
	//determines position of snake head
	x := head[1]

	w.Box(0, 0)

	for _, s := range snake {
		w.MovePrint(head[0], x, s)
	}
	w.Refresh()

}
