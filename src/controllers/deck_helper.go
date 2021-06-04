package controllers

import (
	"caiomcg.com/playing_cards/src/models"
	"errors"
	"strconv"
	"strings"
)

func findDeck(id string) (models.Deck, error) {
	for _, v := range Decks {
		if v.Id.String() == id {
			return v, nil
		}
	}
	return models.Deck{}, errors.New("Deck not found")
}

func processShuffleParam(shuffle string) bool {
	if shuffle == "" {
		shuffle = "false"
	}

	result, _ := strconv.ParseBool(shuffle)
	return result
}

func processCardsParam(cards string) []string {
	if cards == "" {
		return []string{}
	}

	return strings.Split(cards, ",")
}

func getAvailableRange(requested int, available int) int {
	if requested > available {
		return available
	}
	return requested
}
