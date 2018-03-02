package main

func flush(h hand) (bool, []int) {
	var suit rune
	for _, v := range h {
		if suit != 0 && suit != v.Suit {
			return false, []int{h.highCard().Value}
		} else {
			suit = v.Suit
		}
	}
	return true, []int{h.highCard().Value}
}

func straight(h hand) (bool, []int) {
	h.Sort()
	var lastvalue int
	for i, card := range h {
		if i == 0 {
			lastvalue = card.Value
		} else {
			if card.Value == lastvalue+1 {
				lastvalue = card.Value
			} else {
				return false, []int{h.highCard().Value}
			}
		}
	}
	return true, []int{h.highCard().Value}

}

func NOfAKind(h hand, n int) (bool, []int) {
	var occurences map[int]int
	h.Sort()
	occurences = h.Occurences()
	//fmt.Println(occurences)
	for k, value := range occurences {
		if value == n {
			return true, []int{k}
		}
	}
	return false, []int{0}
}

func twoPair(h hand) (bool, []int) {
	var remainingCards hand

	found0, pair0 := NOfAKind(h, 2)

	// set indices for removal
	for _, v := range h {
		if v.Value != pair0[0] {
			remainingCards = append(remainingCards, v)
		}
	}

	found1, pair1 := NOfAKind(remainingCards, 2)

	if found0 && found1 {
		return true, []int{pair0[0], pair1[0]}
	}
	return false, []int{0, 0}
}

func fullHouse(h hand) (bool, []int) {
	var remainingCards hand
	found0, pair0 := NOfAKind(h, 3)

	// set indices for removal
	for _, v := range h {
		if v.Value != pair0[0] {
			remainingCards = append(remainingCards, v)
		}
	}
	found1, pair1 := NOfAKind(remainingCards, 2)

	if found0 && found1 {
		return true, []int{pair0[0], pair1[0]}
	}
	return false, []int{}

}

func straightFlush(h hand) (bool, []int) {
	isFlush, _ := flush(h)
	isStraight, _ := straight(h)
	if isFlush && isStraight {
		return true, []int{h.highCard().Value}
	}
	return false, []int{0}
}

func royalFlush(h hand) (bool, []int) {
	sf, highCard := straightFlush(h)
	if sf && highCard[0] == ace {
		return true, []int{}
	}
	return false, []int{}
}
