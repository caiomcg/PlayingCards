package db

import "testing"
import "caiomcg.com/playing_cards/src/models"
import "fmt"

func createSamples(amount int) []models.Deck {
    decks := []models.Deck{}

    Instance().Wipe()
    for i := 0; i < amount; i++ {
        deck, _ := models.CreateDeck(false, []string{})
        decks = append(decks, deck)
        Instance().Insert(deck)
    }

    return decks
}

func TestInstance(t *testing.T) {
    instance := Instance()

    if instance == nil {
        t.Error("Instance should be a valid pointer")
    }
}

func TestInsert(t *testing.T) {
    decks := createSamples(1)

    deckFromInstance, _ := Instance().Peek()
    if deckFromInstance.Id != decks[0].Id {
        t.Error("Should have the inserted intance")
    }
}

func TestFind(t *testing.T) {
    if _, err := Instance().Find("INVALID"); err == nil {
        t.Error("An error should be raised if no deck is found")
    }

    decks := createSamples(1)

    if content, err := Instance().Find(decks[0].Id.String()); err == nil {
        if content.Id != decks[0].Id {
            t.Error("The deck found should be the same")
        }
    } else {
        t.Error("Should be able to find a deck")
    }
}

func TestGetAll(t *testing.T) {
    createSamples(2)

    if len(Instance().GetAll()) != 2 {
        t.Error("Should have the inseted decks")
    }
}

func TestWipe(t *testing.T) {
    createSamples(10)
    Instance().Wipe()

    if len(Instance().GetAll()) != 0 {
        t.Error("Should not have decks while wiping")
    }
}

func TestPeek(t *testing.T) {
    Instance().Wipe()

    if _, err := Instance().Peek(); err == nil {
        fmt.Println(err)
        t.Error("Should fail if there is nothing to peek")
    }

    decks := createSamples(2)

    if deck, err := Instance().Peek(); err == nil {
        if decks[0].Id != deck.Id {
            t.Error("Should peek the first inserted element")
        }
    } else {
        fmt.Println(err)
        t.Error("Should be able to peek a deck")
    }
}
