package models

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

func (d Deck) Shuffle() {
    shuffleCards(d.Cards)
}

func CreateDeck(shuffle bool, custom []string) Deck {
    var cards []Card
    if len(custom) == 0 {
        cards = GetDefaultSet();
    } else {
        cards = GenerateCustomSet(custom)
    }

    deck := Deck{
        Id: uuid.NewV4(),
        Shuffled: shuffle,
        Remaining: uint8(len(cards)),
        Cards: cards,
    }

    if (shuffle) {
        deck.Shuffle()
    }

    return deck
}

func shuffleCards(cards []Card) {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(cards), func(i, j int) {
        cards[i], cards[j] = cards[j], cards[i]
    })
}
