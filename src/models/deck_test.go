package models

import "testing"

func TestCreateDeck(t *testing.T) {
	regularDeck, _ := CreateDeck(false, []string{})

	if len(regularDeck.Cards) != 52 {
		t.Error("A regular deck should contain 52 cards")
	}

	if regularDeck.Shuffled == true {
		t.Error("The default deck should not be shuffled")
	}
}

func TestShuffle(t *testing.T) {
	regularDeck, _ := CreateDeck(false, []string{})
    regularDeck.Shuffle()

    if regularDeck.Shuffled == false {
        t.Error("The deck should be shuffled")
    }
}

func TestDraw(t *testing.T) {
	regularDeck, _ := CreateDeck(false, []string{})
    cards := regularDeck.Draw(1)

    if regularDeck.Remaining != 51 {
        t.Error("A card should be consumed from the deck")
    }

    if len(cards.Cards) != 1 {
        t.Error("Only one card should be drawn")
    }

    twentyCards := regularDeck.Draw(20)

    if len(twentyCards.Cards) != 20 {
        t.Error("Should have drawn 20 cards")
    }
}

func TestGetAvailableRange(t *testing.T) {
    min := 10
    max := 20

    if getAvailableRange(min, max) != min {
        t.Error("Should get the minimum available range")
    }
}
