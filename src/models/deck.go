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

type DeckDTO struct {
	ID     string `json:"deck_id"`
	Amount int    `json:"amount"`
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
	d.Shuffled = true
}

func CreateDeck(shuffle bool, custom []string) Deck {
	var cards []Card

	if len(custom) == 0 {
		cards = GetDefaultSet()
	} else {
		cards = GenerateCustomSet(custom)
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

	return deck
}
