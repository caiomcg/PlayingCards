package models

import "errors"

type Suits string

const (
	Clubs    Suits = "CLUBS"
	Diamonds Suits = "DIAMONDS"
	Hearts   Suits = "HEARTS"
	Spades   Suits = "SPADES"
)

func (s Suits) GetCode() string {
	return string(s[0])
}

func GetSuitFromCode(code string) (Suits, error) {
	switch code {
	case "C":
		return Clubs, nil
	case "D":
		return Diamonds, nil
	case "H":
		return Hearts, nil
	case "S":
		return Spades, nil
	}

	return "", errors.New("Invalid code")
}
