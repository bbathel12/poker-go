package main

import (
	"math/rand"
	"time"
)

type deck struct {
	cards []card
}

func NewDeck() (d *deck) {
	d = new(deck)

	suits := [4]rune{hearts, diamonds, spades, clubs}
	values := [13]int{two, three, four, five, six, seven, eight, nine,
		ten, jack, queen, king, ace}
	for _, suit := range suits {
		for _, value := range values {
			d.cards = append(d.cards, card{suit, value})
		}
	}
	return
}

func (d *deck) shuffle() {
	for i := 0; i < len(d.cards); i++ {

		rand.Seed(time.Now().UTC().UnixNano())

		r := rand.Intn(len(d.cards))

		d.cards[r], d.cards[i] = d.cards[i], d.cards[r]
	}
}

func (d *deck) pop() (popped card, ok bool) {

	if len(d.cards) >= 1 {
		popped, d.cards = d.cards[0], d.cards[1:]
		ok = true
	} else {
		popped = card{}
		ok = false
	}

	return

}
