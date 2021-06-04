package models

import "testing"

func TestCreateDeck(t *testing.T) {
	regularDeck := CreateDeck(false, []string{})

	if len(regularDeck.Cards) != 52 {
		t.Error("A regular deck should contain 52 cards")
	}

	if regularDeck.Shuffled == true {
		t.Error("The default deck should not be shuffled")
	}
}
