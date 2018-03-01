package main

import (
	_ "fmt"
)

type pokergame struct {
	players  []player
	d        *deck
	handsize int
}

func (pg *pokergame) Setup() {
	pg.players = []player{
		player{"brice", []card{}},
		player{"brice2", []card{}},
	}

	pg.d = NewDeck()
	pg.d.shuffle()
	pg.handsize = 5
}

func (pg *pokergame) Deal() {
	for i := 0; i < pg.handsize; i++ {
		for k, player := range pg.players {
			popped, ok := pg.d.pop()
			if ok {
				pg.players[k].h = append(player.h, popped)
			}
		}
	}
}

func Run() {

	var pg pokergame
	pg.Setup()
	pg.Deal()
	for _, player := range pg.players {
		player.h.Sort()
		player.h.show()
	}

}

func flush(h hand) (flush bool) {
	var suit rune
	flush = true
	for _, v := range h {
		if suit != 0 && suit != v.Suit {
			flush = false
			return
		} else {
			suit = v.Suit
		}
	}
	return
}

func straight(h hand) bool {
	h.Sort()
	var lastvalue int
	for i, card := range h {
		if i == 0 {
			lastvalue = card.Value
		} else {
			if card.Value == lastvalue+1 {
				lastvalue = card.Value
			} else {
				return false
			}
		}
	}
	return true
}

func fourOfAKind(h hand) bool {
	h.Sort()
	var occurences map[int]int = make(map[int]int, 15)

	for _, card := range h {
		occurences[card.Value]++
	}

	for _, value := range occurences {
		if value == 4 {
			return true
		}
	}
	return false
}
