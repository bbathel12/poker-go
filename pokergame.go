package main

import (
	"fmt"
	"sync"
)

const (
	numOfCards = 5
)

type pokergame struct {
	players  []player
	d        *deck
	handsize int
	over     bool
}

func (pg *pokergame) Setup() {
	pg.players = []player{
		player{"brice", []card{}},
		//		player{"brice2", []card{}},
	}

	pg.d = NewDeck()
	pg.d.shuffle()
	pg.handsize = numOfCards
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
	startGame(pg)

	for {
		handleEvent(&pg)
		drawBoard()
		pg.players[0].h.Sort()
		drawHand(pg.players[0].h)
		drawRank(pg.players[0].h)
		drawDiscards()
		draw()
		if pg.over {
			waitForRestart(pg)
			pg.Setup()
			pg.Deal()
			startGame(pg)
			pg.over = false
		}
	}
}

func startGame(pg pokergame) {
	drawBoard()
	for k, _ := range pg.players {
		pg.players[k].h.Sort()
		drawHand(pg.players[k].h)
		drawRank(pg.players[k].h)
		drawDiscards()
	}
	draw()
}

func Simulate() {
	var wg sync.WaitGroup
	var hands map[string]int = map[string]int{}
	var handChan chan string = make(chan string, 1000)

	go func(hands map[string]int) {
		for rank := range handChan {
			hands[rank] += 1
		}
	}(hands)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			var pg pokergame
			for i := 0; i < 1000; i++ {
				pg.Setup()
				pg.Deal()
				for _, player := range pg.players {
					player.h.Sort()
					handChan <- rank(player.h)
					//					player.h.show()
				}
			}
			defer wg.Done()
		}()
	}

	wg.Wait()
	close(handChan)
	for k, v := range hands {
		fmt.Printf("%20s : %4d\n", k, v)
	}
}

func rank(h hand) string {
	if isRoyalFlush, _ := royalFlush(h); isRoyalFlush {
		return "Royal Flush"
	} else if isStraightFlush, _ := straightFlush(h); isStraightFlush {
		return "Straight Flush"
	} else if isFourOfAKind, _ := NOfAKind(h, 4); isFourOfAKind {
		return "Four of a kind"
	} else if isFullHouse, _ := fullHouse(h); isFullHouse {
		return "Full House"
	} else if isFlush, _ := flush(h); isFlush {
		return "Flush"
	} else if isStraight, _ := straight(h); isStraight {
		return "Straight"
	} else if isThreeOfAKind, _ := NOfAKind(h, 3); isThreeOfAKind {
		return "3 of a kind"
	} else if isTwoPair, _ := twoPair(h); isTwoPair {
		return "2 pair"
	} else if isPair, _ := NOfAKind(h, 2); isPair {
		return "Pair"
	} else {
		var cardValue string
		switch h.highCard().Value {
		case 11:
			cardValue = "Jack"
		case 12:
			cardValue = "Queen"
		case 13:
			cardValue = "King"
		case 14:
			cardValue = "Ace"
		}
		return fmt.Sprintf("%s %s", "High Card", cardValue)
	}
}

func removeCards(h hand, discards map[int]bool) hand {
	var remainingCards hand
	remainingCards = make(hand, len(h))
	copy(remainingCards, h)

	for pos, ok := range discards {
		if ok {
			index := pos
			remainingCards[index] = card{}
		}
		discards[pos] = false
	}

	return remainingCards
}

func replaceCards(h hand, d *deck) hand {
	var blankCard card = card{}
	for i, card := range h {
		if card.equal(blankCard) {
			if newCard, ok := d.pop(); ok {
				h[i] = newCard
			}
		}
	}
	return h
}
