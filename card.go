package main

import (
	_ "fmt"
)

const (
	spades   = '♠'
	hearts   = '♥'
	diamonds = '♦'
	clubs    = '♣'
	height   = 20
	width    = 15
)

const (
	_ = iota
	_
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
)

type corner struct {
	X, Y int
}

type pixel struct {
	X, Y   int
	Symbol rune
}

type card struct {
	Suit  rune
	Value int
}

func (c card) draw(topLeft corner) []pixel {
	var startx int = topLeft.X
	var endx int = topLeft.X + width
	var starty int = topLeft.Y
	var endy int = topLeft.Y + height
	var pix pixel
	var pixels []pixel
	for x := startx; x <= endx; x++ {
		for y := starty; y <= endy; y++ {
			pix = pixel{x, y, ' '}
			pixels = append(pixels, pix)
		}
	}
	// SUIT SPOTS
	pixels = append(pixels, pixel{startx + 2, starty + 2, c.Suit})
	pixels = append(pixels, pixel{endx - 2, starty + 2, c.Suit})
	pixels = append(pixels, pixel{startx + 2, endy - 2, c.Suit})
	pixels = append(pixels, pixel{endx - 2, endy - 2, c.Suit})

	// number SPOTS
	//pixels = append(pixels, pixel{startx + 1, starty + 1, c.Value})
	//pixels = append(pixels, pixel{endx - 1, starty + 1, c.Value})
	//pixels = append(pixels, pixel{startx + 1, endy - 1, c.Value})
	//pixels = append(pixels, pixel{endx - 1, endy - 1, c.Value})

	return pixels
}

func (c card) equal(c2 card) bool {
	if c.Suit == c2.Suit && c.Value == c.Value {
		return true
	}
	return false
}
