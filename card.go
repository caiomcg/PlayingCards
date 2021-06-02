package main

import "fmt"

// TODO: MAP 10?

type Card struct {
    Number string `json:"value"`
    Suit Suits `json:"suit"`
    Code string `json:"code"`
}

func MapNumber(value uint8) string {
    if value == 1 {
        return "ACE"
    }

    if value == 10 {
        return "0"
    }

    if value == 11 {
        return "JACK"
    }

    if value == 12 {
        return "QUEEN"
    }

    if value == 13 {
        return "KING"
    }

    return fmt.Sprint(value);
}

func GetCardCode(name string, suit Suits) string {
    suit_code := suit.GetCode()
    number_code := string(name[0])

    return number_code + suit_code
}

func CreateCardWith(code string) {
    fmt.Println("ERRO: Need to create a card from its code")
}

func CreateCard(number uint8, suit Suits) Card {
    num := MapNumber(number)

    return Card{Number: num, Suit: suit, Code: GetCardCode(num, suit)}
}

func GenerateSet(suit Suits) []Card {
    cards := []Card{}

    for i := 1; i < 14; i++ {
        cards = append(cards, CreateCard(uint8(i), suit));
    }

    return cards
}

func GetDefaultSet() []Card {
    cards := []Card{}

    cards = append(cards, GenerateSet(Spades)...)
    cards = append(cards, GenerateSet(Diamonds)...)
    cards = append(cards, GenerateSet(Clubs)...)
    cards = append(cards, GenerateSet(Hearts)...)

    return cards;
}

