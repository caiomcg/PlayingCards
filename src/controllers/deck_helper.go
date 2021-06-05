package controllers

import (
	"errors"
	"strconv"
	"strings"
)

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
