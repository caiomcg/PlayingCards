package controllers

import (
	"caiomcg.com/playing_cards/src/helpers"
	"caiomcg.com/playing_cards/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

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
	deckDTO := new(models.DeckDTO)

	if err := c.Bind(deckDTO); err != nil {
		return helpers.NewHTTPError(
			http.StatusInternalServerError,
			"ParserError",
			"Could not parse the given parameters",
		)
	}

	if deckDTO.ID == "" || deckDTO.Amount == 0 {
		return helpers.NewHTTPError(
			http.StatusBadRequest,
			"InvalidParams",
			"Both deck_id and amount have to be valid in order to process the content",
		)
	}

	deck, e := findDeck(deckDTO.ID)

	if e != nil {
		return helpers.NewHTTPError(
			http.StatusNotFound,
			"Invalid deck_id",
			"Could not find a deck with the desired ID",
		)
	}

	cards := models.Cards{
		Cards: deck.Cards[0:getAvailableRange(int(deckDTO.Amount), len(deck.Cards))],
	}

	return c.JSON(
		http.StatusOK,
		cards,
	)
}

