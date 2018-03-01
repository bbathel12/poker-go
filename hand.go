package main

import (
	"fmt"
	"sort"
	"strconv"
)

type hand []card

func (h hand) Len() int {
	return len(h)
}

func (h hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hand) Less(i, j int) bool {
	return h[i].Value < h[j].Value
}

func (h *hand) Sort() {
	sort.Sort(h)
}

func (h hand) show() {
	for _, v := range h {
		var value string
		switch v.Value {
		case 14:
			value = "A"
		case 13:
			value = "K"
		case 12:
			value = "Q"
		case 11:
			value = "J"
		default:
			value = strconv.Itoa(v.Value)
		}

		fmt.Printf("%2s%1c ", value, v.Suit)
	}
	fmt.Println()
}
