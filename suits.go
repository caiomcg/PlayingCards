package main

type Suits string

const (
    Clubs Suits = "CLUBS"
    Diamonds = "DIAMONDS"
    Hearts = "HEARTS"
    Spades = "SPADES"
)

func (s Suits) GetCode() string {
    return string(s[0]);
}

