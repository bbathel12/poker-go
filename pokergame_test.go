package main

import (
	"testing"
)

func Test_Setup(t *testing.T) {

	var pg pokergame
	pg.Setup()

	if len(pg.d.cards) == 0 {
		t.Error("deck is empty slice")
	}

	if pg.handsize == 0 {
		t.Error("handsize is 0")
	}
	if len(pg.players) == 0 {
		t.Error("players is empty")
	}
}

func Test_flush(t *testing.T) {
	var hands []hand = []hand{
		hand{
			card{hearts, 10},
			card{hearts, 9},
			card{hearts, 4},
			card{hearts, 14},
			card{hearts, 12},
		},
		hand{
			card{hearts, 10},
			card{hearts, 9},
			card{spades, 4},
			card{hearts, 14},
			card{hearts, 12},
		},
	}

	if !flush(hands[0]) {
		hands[0].show()
		t.Error("should be flush")
	}

	if flush(hands[1]) {
		hands[1].show()
		t.Error("Should not be flush")
	}
}

func Test_Straight(t *testing.T) {
	var hands []hand = []hand{
		hand{
			card{hearts, 10},
			card{hearts, 9},
			card{hearts, 8},
			card{hearts, 7},
			card{hearts, 6},
		},
		hand{
			card{hearts, 10},
			card{hearts, 9},
			card{spades, 4},
			card{hearts, 14},
			card{hearts, 12},
		},
	}

	if !straight(hands[0]) {
		hands[0].show()
		t.Error("should be straight")
	}

	if straight(hands[1]) {
		hands[1].show()
		t.Error("Should not be straight")
	}
}

func Test_fourOfAKind(t *testing.T) {
	var hands []hand = []hand{
		hand{
			card{hearts, 10},
			card{spades, 10},
			card{diamonds, 10},
			card{clubs, 10},
			card{hearts, 6},
		},
		hand{
			card{hearts, 10},
			card{hearts, 9},
			card{spades, 4},
			card{hearts, 14},
			card{hearts, 12},
		},
	}

	if !fourOfAKind(hands[0]) {
		hands[0].show()
		t.Error("should be 4 of a kind")
	}

	if fourOfAKind(hands[1]) {
		hands[1].show()
		t.Error("Should not be 4 of a kind")
	}
}
