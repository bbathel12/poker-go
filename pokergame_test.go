package main

import (
	"fmt"
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

	isFlush, _ := flush(hands[0])
	if !isFlush {
		hands[0].show()
		t.Error("should be flush")
	}

	isFlush, _ = flush(hands[1])
	if isFlush {
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

	isStraight, _ := straight(hands[0])
	if !isStraight {
		hands[0].show()
		t.Error("should be straight")
	}

	isStraight, _ = straight(hands[1])
	if isStraight {
		hands[1].show()
		t.Error("Should not be straight")
	}
}

func Test_NOfAKind(t *testing.T) {
	var hands []hand = []hand{
		hand{ //4 of a kind
			card{hearts, 10},
			card{spades, 10},
			card{diamonds, 10},
			card{clubs, 10},
			card{hearts, 6},
		},
		hand{ // no copies
			card{hearts, 10},
			card{hearts, 9},
			card{spades, 4},
			card{hearts, 14},
			card{hearts, 12},
		},
		hand{ // 3 of a kind
			card{hearts, 4},
			card{hearts, 3},
			card{spades, 4},
			card{clubs, 4},
			card{spades, 9},
		},
		hand{
			card{hearts, 2},
			card{clubs, 2},
			card{spades, 4},
			card{spades, 10},
			card{spades, 11},
		},
	}
	ok, card := NOfAKind(hands[0], 4)
	if !ok || card[0] != 10 {
		hands[0].show()
		fmt.Println(ok, card)
		t.Error("should be 4 of a kind")
	}

	ok, card = NOfAKind(hands[1], 4)
	if ok || card[0] != 0 {
		hands[1].show()
		fmt.Println(ok, card)
		t.Error("Should not be 4 of a kind")
	}

	ok, card = NOfAKind(hands[2], 3)
	if !ok || card[0] != 4 {
		hands[2].show()
		fmt.Println(ok, card)
		t.Errorf("Should not be %d of a kind", 3)
	}

	ok, card = NOfAKind(hands[3], 2)
	if !ok || card[0] != 2 {
		hands[3].show()
		fmt.Println(ok, card)
		t.Errorf("Should not be %d of a kind", 2)
	}

}

func Test_twoPair(t *testing.T) {
	var hands []hand = []hand{
		hand{
			card{hearts, 2},
			card{hearts, 3},
			card{spades, 2},
			card{spades, 3},
			card{spades, 8},
		},
		hand{
			card{hearts, 2},
			card{hearts, 3},
			card{hearts, 4},
			card{hearts, 5},
			card{hearts, 6},
		},
	}

	ok, cards := twoPair(hands[0])
	if !ok {
		fmt.Println(ok, cards)
		t.Error("should be two pair 2's and 3's")
	}

	ok, cards = twoPair(hands[1])
	if ok || cards[0] != 0 || cards[1] != 0 {
		fmt.Println(ok, cards)
		t.Error("should not be two pair")
	}

}

func Test_fullHouse(t *testing.T) {
	var hands []hand = []hand{
		hand{
			card{hearts, 2},
			card{hearts, 2},
			card{spades, 2},
			card{spades, 3},
			card{spades, 3},
		},
		hand{
			card{hearts, 2},
			card{hearts, 3},
			card{hearts, 4},
			card{hearts, 5},
			card{hearts, 6},
		},
	}

	ok, cards := fullHouse(hands[0])
	if !ok || cards[0] != 2 || cards[1] != 3 {
		fmt.Println(ok, cards)
		t.Error("should be fullhouse  2's and 3's")
	}

	ok, cards = twoPair(hands[1])
	if ok || cards[0] != 0 || cards[1] != 0 {
		fmt.Println(ok, cards)
		t.Error("should not be two pair")
	}

}

func Test_straightFlush(t *testing.T) {
	var hands []hand = []hand{
		hand{
			card{hearts, 2},
			card{hearts, 3},
			card{hearts, 4},
			card{hearts, 5},
			card{hearts, 6},
		},
		hand{
			card{hearts, 2},
			card{hearts, 3},
			card{hearts, 8},
			card{hearts, 5},
			card{hearts, 6},
		},
	}

	ok, highcard := straightFlush(hands[0])
	if !ok || highcard[0] != 6 {
		fmt.Println(ok, highcard)
		t.Error("should be straight flush with high card 6")
	}

	ok, highcard = straightFlush(hands[1])
	if ok || highcard[0] != 0 {
		fmt.Println(ok, highcard)
		t.Error("should not be straigh flush")
	}

}

func Test_royalFlush(t *testing.T) {
	var hands []hand = []hand{
		hand{
			card{hearts, ace},
			card{hearts, king},
			card{hearts, queen},
			card{hearts, jack},
			card{hearts, ten},
		},
		hand{
			card{hearts, 2},
			card{hearts, 3},
			card{hearts, 8},
			card{hearts, 5},
			card{hearts, 6},
		},
	}
	isRoyalFlush, _ := royalFlush(hands[0])
	if !isRoyalFlush {
		t.Error("should be royal flush ")
	}

	isRoyalFlush, _ = royalFlush(hands[1])
	if isRoyalFlush {
		t.Error("should not be royal flush")
	}

}
