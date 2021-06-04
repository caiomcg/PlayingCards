package models

import "testing"

func showError(t *testing.T, suit string, response string, expected string) {
	t.Errorf("Test for %s returned %s instead of %s", suit, response, expected)
}

func TestGetCode(t *testing.T) {
	clubs := Clubs
	diamonds := Diamonds
	hearts := Hearts
	spades := Spades

	if clubs.GetCode() != "C" {
		showError(t, "clubs", clubs.GetCode(), "C")
	}
	if diamonds.GetCode() != "D" {
		showError(t, "diamonds", diamonds.GetCode(), "D")
	}
	if hearts.GetCode() != "H" {
		showError(t, "hearts", hearts.GetCode(), "H")
	}
	if spades.GetCode() != "S" {
		showError(t, "spades", spades.GetCode(), "S")
	}
}

func testForCode(t *testing.T, code string) {
	if _, err := GetSuitFromCode(code); err != nil {
		t.Error("Should be able to get ")
	}
}

func TestGetSuitFromCode(t *testing.T) {
	clubs := "C"
	diamonds := "D"
	hearts := "H"
	spades := "S"

	testForCode(t, clubs)
	testForCode(t, diamonds)
	testForCode(t, hearts)
	testForCode(t, spades)

	if _, err := GetSuitFromCode("random"); err == nil {
		t.Error("Should not return a code for a random suit")
	}
}
