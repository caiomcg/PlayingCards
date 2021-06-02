package main

import (
	"math/rand"
	"time"
	"github.com/satori/go.uuid"
)

type Deck struct {
    Id uuid.UUID `json:"deck_id"`
    Shuffled bool `json:"shuffled"`
    Remaining uint8 `json:"remaining"`
    Cards []Card `json:"cards"`
}

func ShuffleCards(cards []Card) {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(cards), func(i, j int) {
        cards[i], cards[j] = cards[j], cards[i]
    })
}

func CreateDeck(shuffle bool) Deck {
    cards := GetDefaultSet();

    if shuffle {
        ShuffleCards(cards)
    }

    return Deck{
        Id: uuid.NewV4(),
        Shuffled: shuffle,
        Remaining: 52,
        Cards: cards,
    }
}
