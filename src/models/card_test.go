package models

import "testing"

var codes = []struct {
	name   string
	suit   Suits
	code   string
	number uint8
}{
	{"ACE", Spades, "AS", 1},
	{"ACE", Clubs, "AC", 1},
	{"ACE", Diamonds, "AD", 1},
	{"ACE", Hearts, "AH", 1},
	{"2", Spades, "2S", 2},
	{"2", Clubs, "2C", 2},
	{"2", Diamonds, "2D", 2},
	{"2", Hearts, "2H", 2},
	{"3", Spades, "3S", 3},
	{"3", Clubs, "3C", 3},
	{"3", Diamonds, "3D", 3},
	{"3", Hearts, "3H", 3},
	{"4", Spades, "4S", 4},
	{"4", Clubs, "4C", 4},
	{"4", Diamonds, "4D", 4},
	{"4", Hearts, "4H", 4},
	{"5", Spades, "5S", 5},
	{"5", Clubs, "5C", 5},
	{"5", Diamonds, "5D", 5},
	{"5", Hearts, "5H", 5},
	{"6", Spades, "6S", 6},
	{"6", Clubs, "6C", 6},
	{"6", Diamonds, "6D", 6},
	{"6", Hearts, "6H", 6},
	{"7", Spades, "7S", 7},
	{"7", Clubs, "7C", 7},
	{"7", Diamonds, "7D", 7},
	{"7", Hearts, "7H", 7},
	{"8", Spades, "8S", 8},
	{"8", Clubs, "8C", 8},
	{"8", Diamonds, "8D", 8},
	{"8", Hearts, "8H", 8},
	{"9", Spades, "9S", 9},
	{"9", Clubs, "9C", 9},
	{"9", Diamonds, "9D", 9},
	{"9", Hearts, "9H", 9},
	{"X", Spades, "XS", 10},
	{"X", Clubs, "XC", 10},
	{"X", Diamonds, "XD", 10},
	{"X", Hearts, "XH", 10},
	{"JACK", Spades, "JS", 11},
	{"JACK", Clubs, "JC", 11},
	{"JACK", Diamonds, "JD", 11},
	{"JACK", Hearts, "JH", 11},
	{"QUEEN", Spades, "QS", 12},
	{"QUEEN", Clubs, "QC", 12},
	{"QUEEN", Diamonds, "QD", 12},
	{"QUEEN", Hearts, "QH", 12},
	{"KING", Spades, "KS", 13},
	{"KING", Clubs, "KC", 13},
	{"KING", Diamonds, "KD", 13},
	{"KING", Hearts, "KH", 13},
}

var invalidCodes = []string{
	"BS", "ZC", "DD", "SH",
	"%S", "*C", "!D", "0H",
	"A2", "AZ", "A!", "A$",
}

func TestIsCodeValid(t *testing.T) {
	for _, tt := range codes {
		if !isCodeValid(tt.code) {
			t.Errorf("Should be valid for %s", tt.code)
		}
	}

	for _, code := range invalidCodes {
		if isCodeValid(code) {
			t.Errorf("Should be invalid for %s", code)
		}
	}
}

func TestMapNumber(t *testing.T) {
	for _, tt := range codes {
		if mapNumber(tt.number) != tt.name {
			t.Errorf("%d should map to %s", tt.number, tt.name)
		}
	}
}

func TestMapCode(t *testing.T) {
	var numbers = []struct {
		in  string
		out uint8
	}{
		{"A", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
		{"X", 10},
		{"J", 11},
		{"Q", 12},
		{"K", 13},
	}

	for _, tt := range numbers {
		if mapCode(tt.in) != tt.out {
			t.Errorf("%s should map to %d", tt.in, tt.out)
		}
	}
}

func TestGetCardCode(t *testing.T) {
	for _, tt := range codes {
		if getCardCode(tt.name, tt.suit) != tt.code {
			t.Errorf(
				"Should get %s from a combination of %s and %s",
				tt.code,
				tt.name,
				tt.suit,
			)
		}
	}
}

func compareSet(t *testing.T, suit Suits, sample []Card) {
	if len(sample) != 13 {
		t.Error("A default set should contain 13 cards")
	}

	for _, v := range sample {
		if v.Suit != suit {
			t.Error("A set should have only one suit")
		}
	}
}

func TestGenerateSet(t *testing.T) {
	compareSet(t, Spades, generateSet(Spades))
	compareSet(t, Diamonds, generateSet(Diamonds))
	compareSet(t, Clubs, generateSet(Clubs))
	compareSet(t, Hearts, generateSet(Hearts))
}

func TestCreateCardFromCode(t *testing.T) {
	if _, err := CreateCardFromCode("INVALID"); err == nil {
		t.Error("Should fail if an invalid code is given")
	}

	for _, tt := range codes {
		if _, err := CreateCardFromCode(tt.code); err != nil {
			t.Errorf("Should create a card from a valid code: %s", tt.code)
		}
	}
}

func TestGetDefaultSet(t *testing.T) {
	if len(GetDefaultSet()) != 52 {
		t.Error("A default set should contain 52 cards")
	}
}

func deepCompare(cards []Card, codes []string) bool {
	if len(cards) != len(codes) {
		return false
	}

	for i := 0; i < len(cards); i++ {
		if cards[i].Code != codes[i] {
			return false
		}
	}

	return true
}
func TestGenerateCustomSet(t *testing.T) {
	emptySet, _ := GenerateCustomSet([]string{})

	if len(emptySet) != 0 {
		t.Error("Should create an empty deck")
	}

	oneCardCode := []string{"AS"}
	oneCardSet, _ := GenerateCustomSet(oneCardCode)
	if len(oneCardSet) != 1 || !deepCompare(oneCardSet, oneCardCode) {
		t.Error("Should create a deck with only one card")
	}

	multipleCardCode := []string{"AS", "2H", "3D", "KC"}
	multipleCardDeck, _ := GenerateCustomSet(multipleCardCode)
	if len(multipleCardDeck) < 2 || !deepCompare(multipleCardDeck, multipleCardCode) {
		t.Error("Should create a set with multiple cards")
	}

	repeatedCardCode := []string{"AS", "AS", "3D", "KC"}
	repeatedCardDeck, _ := GenerateCustomSet(repeatedCardCode)
	if len(repeatedCardDeck) != 3 || deepCompare(repeatedCardDeck, repeatedCardCode) {
		t.Error("Should create a set with no duplicate cards")
	}

	invalidCardCode := []string{"ZZ"}
	_, e := GenerateCustomSet(invalidCardCode)
	if e == nil {
		t.Error("Should fail if the code is invalid")
	}
}
