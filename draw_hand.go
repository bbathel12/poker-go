package main

import (
	_ "fmt"
	"github.com/nsf/termbox-go"
	"os"
	"time"
)

var cardWidth int = 5
var width int
var discards map[int]bool = map[int]bool{0: false, 1: false, 2: false, 3: false, 4: false}

func init() {
	termbox.Init()
	width, _ = termbox.Size()

}

func drawHand(h hand) {
	for i, card := range h {
		drawCard(card, i)
	}
}

func drawBoard() {
	for i := 0; i <= width; i++ {
		termbox.SetCell(i, 1, 0, termbox.ColorBlue, termbox.ColorYellow)
		termbox.SetCell(i, 2, 0, termbox.ColorBlue, termbox.ColorYellow)
		termbox.SetCell(i, 3, 0, termbox.ColorBlue, termbox.ColorYellow)
		termbox.SetCell(i, 4, 0, termbox.ColorBlue, termbox.ColorYellow)
	}
}

func drawCard(c card, position int) {
	var suit [2]rune
	position = position * cardWidth
	suit = convertToRunes(c)
	position += 2
	termbox.SetCell(position, 1, suit[0], termbox.ColorBlue, termbox.ColorYellow)
	position++
	termbox.SetCell(position, 1, suit[1], termbox.ColorBlue, termbox.ColorYellow)
	position++
	termbox.SetCell(position, 1, c.Suit, termbox.ColorBlue, termbox.ColorYellow)

}

func drawRank(h hand) {
	position := cardWidth * 5
	position += 2
	hrank := []rune(rank(h))
	for _, char := range hrank {
		position++

		termbox.SetCell(position, 1, char, termbox.ColorBlue, termbox.ColorYellow)
	}
}

func convertToRunes(c card) (suit [2]rune) {
	if c.Value == 10 {
		suit[0] = 49
		suit[1] = 48
	} else if c.Value > 10 {
		switch c.Value {
		case 11:
			suit[1] = 74
		case 12:
			suit[1] = 81
		case 13:
			suit[1] = 75
		case 14:
			suit[1] = 65
		default:

		}
	} else {
		suit[0] = 0
		switch c.Value {
		case 9:
			suit[1] = 57
		case 8:
			suit[1] = 56
		case 7:
			suit[1] = 55
		case 6:
			suit[1] = 54
		case 5:
			suit[1] = 53
		case 4:
			suit[1] = 52
		case 3:
			suit[1] = 51
		case 2:
			suit[1] = 50
		}
	}
	return
}
func drawDiscards() {
	for k, v := range discards {
		if !v {
			position := k * cardWidth
			position++
			termbox.SetCell(position, 3, 'K', termbox.ColorBlue, termbox.ColorYellow)
			position++
			termbox.SetCell(position, 3, 'E', termbox.ColorBlue, termbox.ColorYellow)
			position++
			termbox.SetCell(position, 3, 'E', termbox.ColorBlue, termbox.ColorYellow)
			position++
			termbox.SetCell(position, 3, 'P', termbox.ColorBlue, termbox.ColorYellow)
		}
	}
}

func handleEvent(pg *pokergame) {
	e := termbox.PollEvent()
	if e.Key == termbox.KeyEnd {
		os.Exit(1)
	}

	switch e.Key {
	case termbox.KeyEnd, termbox.KeyBackspace, termbox.KeyBackspace2:
		os.Exit(1)
	case 13:
		pg.players[0].h = removeCards(pg.players[0].h, discards)
		pg.players[0].h = replaceCards(pg.players[0].h, pg.d)
		pg.over = true

	}

	switch e.Ch {
	case '1':
		discards[0] = !discards[0]
		//showDelete(1)
		//		fmt.Println("you pressed 1")
	case '2':
		discards[1] = !discards[1]
		//showDelete(2)
		//		fmt.Println("you pressed 2")
	case '3':
		discards[2] = !discards[2]
		//		fmt.Println("you pressed 3")
		//showDelete(3)
	case '4':
		discards[3] = !discards[3]
		//		fmt.Println("you pressed 4")
		// showDelete(4)
	case '5':
		discards[4] = !discards[4]
		//		fmt.Println("you pressed 5")
		//showDelete(5)
	}
}

func draw() {
	drawInstructions()
	termbox.Flush()
}
func waitForRestart(pg pokergame) {
	i := 0
	var kill chan bool = make(chan bool)
	go func(i int, pg pokergame) {
	gameoverloop:
		for {
			select {
			case <-kill:
				break gameoverloop
			default:
				drawBoard()
				pg.players[0].h.Sort()
				drawHand(pg.players[0].h)
				drawRank(pg.players[0].h)
				i++
				drawGameOver(i)
				draw()
			}
			time.Sleep(time.Millisecond * 500)
		}
	}(i, pg)
loop:
	for {
		e := termbox.PollEvent()
		switch e.Key {
		case 13:
			kill <- true
			close(kill)
			break loop
		case termbox.KeyEnd, termbox.KeyBackspace2, termbox.KeyBackspace:
			os.Exit(1)

		}
	}
}
func drawGameOver(i int) {
	var beg, end string
	switch i % 6 {
	case 0:
		beg = "    ☄☄☄⚀"
		end = "⚀☄☄☄    "
	case 1:
		beg = "    ☣☣☣⚁"
		end = "⚁☣☣☣    "
	case 2:
		beg = "     ⚞⚞⚞⚂"
		end = "⚂⚟⚟⚟     "
	case 3:
		beg = "    ⚡⚡⚡ ⚃"
		end = "⚃ ⚡⚡⚡    "
	case 4:
		beg = "   ⛏⛏⛏ ⚄"
		end = "⚄ ⛏⛏⛏   "
	case 5:
		beg = "      ⛮⛮⛮ ⚅"
		end = "⚅ ⛮⛮⛮      "
	}
	str := beg + "Game Over" + end

	for k, v := range []rune(str) {
		termbox.SetCell(k, 4, v, termbox.ColorMagenta, termbox.ColorCyan)
	}
}

func drawInstructions() {
	instructions := "-Press Backspace or end to quit at anytime"
	instructions2 := "-Press number keys to select cards to discard(1-5)"
	instructions3 := "-Press enter to draw new cards"
	instructions4 := "-Press enter to start again after game over"
	for i := 0; i < width; i++ {
		termbox.SetCell(i, 8, 0, termbox.ColorRed, termbox.ColorBlack)
		termbox.SetCell(i, 9, ' ', termbox.ColorRed, termbox.ColorBlack)
		termbox.SetCell(i, 10, ' ', termbox.ColorRed, termbox.ColorBlack)
		termbox.SetCell(i, 11, ' ', termbox.ColorRed, termbox.ColorBlack)
	}
	for k, v := range []rune(instructions) {
		termbox.SetCell(k, 8, v, termbox.ColorRed, termbox.ColorBlack)
	}
	for k, v := range []rune(instructions2) {
		termbox.SetCell(k, 9, v, termbox.ColorRed, termbox.ColorBlack)
	}
	for k, v := range []rune(instructions3) {
		termbox.SetCell(k, 10, v, termbox.ColorRed, termbox.ColorBlack)
	}
	for k, v := range []rune(instructions4) {
		termbox.SetCell(k, 11, v, termbox.ColorRed, termbox.ColorBlack)
	}
}
