package models

import (
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

type Deck struct {
	Id        uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining uint8     `json:"remaining"`
	Cards     []Card    `json:"cards"`
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
	d.Shuffled = true
}

func (d *Deck) Draw(count int) Cards {
	amountToDraw := getAvailableRange(count, len(d.Cards))

	cards := Cards{
		Cards: d.Cards[0:amountToDraw],
	}

	d.Cards = d.Cards[amountToDraw:]
	d.Remaining = d.Remaining - uint8(amountToDraw)

	return cards
}

func CreateDeck(shuffle bool, custom []string) (Deck, error) {
	var cards []Card

	if len(custom) == 0 {
		cards = GetDefaultSet()
	} else {
		customCards, e := GenerateCustomSet(custom)
		if e != nil {
			return Deck{}, e
		}
		cards = customCards
	}

	deck := Deck{
		Id:        uuid.NewV4(),
		Shuffled:  false,
		Remaining: uint8(len(cards)),
		Cards:     cards,
	}

	if shuffle {
		deck.Shuffle()
	}

	return deck, nil
}

func getAvailableRange(requested int, available int) int {
	if requested > available {
		return available
	}
	return requested
}
