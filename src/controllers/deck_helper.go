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

func processAmountParam(amount string) (int64, error) {
	if amount == "" {
		return 0, errors.New("An amount is required")
	}

	amount_num, err := strconv.ParseInt(amount, 10, 64)

	if err != nil || amount_num <= 0 {
		return 0, errors.New("Invalid amount")
	}

	return amount_num, nil
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
