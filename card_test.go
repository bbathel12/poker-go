package main

import (
	_ "fmt"
	"testing"
)

//func Test_draw(t *testing.T) {
//	var testCorners []corner
//	var testCards []card
//	testCorners = []corner{
//		corner{10, 10},
//		corner{1, 2},
//		corner{5, 9},
//		corner{-1, -1},
//		corner{1, 7},
//	}
//
//	testCards = []card{
//		card{spades, 10},
//		card{hearts, 3},
//		card{diamonds, 1},
//		card{clubs, 7},
//		card{clubs, -1},
//	}
//
//	for k, card := range testCards {
//		_ = card.draw(testCorners[k])
//		//fmt.Println(pixels)
//	}
//}

func Test_equal(t *testing.T) {
	h10 := card{hearts, 10}
	d10 := card{diamonds, 10}
	//	s10 := card{spades, 10}
	c1 := card{clubs, 1}
	c1c := card{clubs, 1}

	if h10.equal(d10) {
		t.Error("h10 shouldn't equal d10")
	}

	if d10.equal(c1) {
		t.Error("d10 shouldn't equal c1")
	}

	if !c1.equal(c1c) {
		t.Error("c1 should equal c1c")
	}
}

func Test_shuffle(t *testing.T) {
	d := NewDeck()
	copyofcards := make([]card, len(d.cards))
	copy(copyofcards, d.cards)
	d.shuffle()

	var matches int = 0
	for k, card := range d.cards {
		if card.equal(copyofcards[k]) {
			matches++
		}
	}
	//	fmt.Println(matches)
	if matches > len(d.cards)/2 {
		t.Error("more than 50% match")
	}

}

func Test_pop(t *testing.T) {
	d := NewDeck()
	d.shuffle()

	for _, card := range d.cards {
		oldLength := len(d.cards)
		popped, ok := d.pop()
		if ok && !popped.equal(card) {
			t.Error(popped, "does not equal", card)
		}
		if oldLength == len(d.cards)-1 {
			t.Error(oldLength, "didn't shrink", len(d.cards))
		}
	}
}

func Benchmark_shuffle(b *testing.B) {
	d := NewDeck()

	for i := 0; i <= b.N; i++ {
		d.shuffle()
	}

}
