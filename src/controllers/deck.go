package controllers

import (
	"caiomcg.com/playing_cards/src/helpers"
	"caiomcg.com/playing_cards/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// This is our "DB". To avoid taking more time studying an in memory database
// I've decided upon using a simple array as our database. This should be
// moved to a dedicated DB if we care about ACID
var Decks []models.Deck = []models.Deck{}

func CreateDeckEndpoint(c echo.Context) error {
	shuffle := processShuffleParam(c.QueryParam("shuffle"))
	deck := processCardsParam(c.QueryParam("cards"))

	newDeck := models.CreateDeck(shuffle, deck)

	Decks = append(Decks, newDeck)
	return c.JSON(http.StatusOK, newDeck)
}

func FetchDecksEndpoint(c echo.Context) error {
	return c.JSON(http.StatusOK, Decks)
}

func OpenDeckEndpoint(c echo.Context) error {
	id := c.Param("id")

	deck, e := findDeck(id)

	if e != nil {
		return helpers.NewHTTPError(
			http.StatusNotFound,
			"Invalid deck_did",
			"Could not find a deck with the desired ID",
		)
	}

	return c.JSON(http.StatusOK, deck)
}

func FetchDeckCardsEndpoint(c echo.Context) error {
	deck, e := findDeck(c.QueryParam("id"))
	if e != nil {
		return helpers.NewHTTPError(
			http.StatusNotFound,
			"Invalid id",
			"Could not find a deck with the desired ID",
		)
	}

	amount, e := processAmountParam(c.QueryParam("amount"))
	if e != nil {
		return helpers.NewHTTPError(
			http.StatusBadRequest,
			"Invalid amount",
			"Amount not given or zero",
		)
	}
	cards := models.Cards{
		Cards: deck.Cards[0:getAvailableRange(int(amount), len(deck.Cards))],
	}

	return c.JSON(
		http.StatusOK,
		cards,
	)
}
