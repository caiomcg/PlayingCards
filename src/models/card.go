package models

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Card struct {
	Number string `json:"value"`
	Suit   Suits  `json:"suit"`
	Code   string `json:"code"`
}

func isCodeValid(code string) bool {
	regex, _ := regexp.Compile("^([A|J|K|Q|X]|[1-9])([S|D|C|H])$")

	return regex.MatchString(code)
}

func mapNumber(value uint8) string {
	switch value {
	case 1:
		return "ACE"
	case 10:
		return "X"
	case 11:
		return "JACK"
	case 12:
		return "QUEEN"
	case 13:
		return "KING"
	}

	return fmt.Sprint(value)
}

func mapCode(code string) uint8 {
	switch code {
	case "A":
		return 1
	case "X":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	}

	num, _ := strconv.ParseUint(code, 10, 8)
	return uint8(num)
}

func getCardCode(name string, suit Suits) string {
	suit_code := suit.GetCode()
	number_code := string(name[0])

	return number_code + suit_code
}

func generateSet(suit Suits) []Card {
	cards := []Card{}

	for i := 1; i < 14; i++ {
		cards = append(cards, CreateCard(uint8(i), suit))
	}

	return cards
}

func CreateCardFromCode(code string) (Card, error) {
	if !isCodeValid(code) {
		return Card{}, errors.New("Invalid code")
	}

	num_code := mapCode(string(code[0]))
	num := mapNumber(num_code)
	suit, _ := GetSuitFromCode(string(code[1]))

	return Card{Number: num, Suit: suit, Code: getCardCode(num, suit)}, nil
}

func CreateCard(number uint8, suit Suits) Card {
	num := mapNumber(number)

	return Card{Number: num, Suit: suit, Code: getCardCode(num, suit)}
}

func GetDefaultSet() []Card {
	cards := []Card{}

	cards = append(cards, generateSet(Spades)...)
	cards = append(cards, generateSet(Diamonds)...)
	cards = append(cards, generateSet(Clubs)...)
	cards = append(cards, generateSet(Hearts)...)

	return cards
}

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {

		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

func GenerateCustomSet(in []string) ([]Card, error) {
	result := []Card{}

	for _, v := range removeDuplicates(in) {
		card, e := CreateCardFromCode(v)
		if e != nil {
			return []Card{}, e
		}
		result = append(result, card)
	}

	return result, nil
}
