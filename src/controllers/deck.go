package controllers

import (
	"caiomcg.com/playing_cards/src/db"
	"caiomcg.com/playing_cards/src/helpers"
	"caiomcg.com/playing_cards/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// This is our "DB". To avoid taking more time studying an in memory database
// I've decided upon using a simple array as our database. This should be
// moved to a dedicated DB if we care about ACID
func CreateDeckEndpoint(c echo.Context) error {
	shuffle := processShuffleParam(c.QueryParam("shuffle"))
	deck := processCardsParam(c.QueryParam("cards"))

	newDeck, e := models.CreateDeck(shuffle, deck)
	if e != nil {
		return helpers.NewHTTPError(
			http.StatusBadRequest,
			"Invalid custom card",
			"Card codes should follow the estipulated rules",
		)
	}

	db.Instance().Insert(newDeck)
	return c.JSON(http.StatusOK, newDeck)
}

func FetchDecksEndpoint(c echo.Context) error {
	return c.JSON(http.StatusOK, db.Instance().GetAll())
}

func OpenDeckEndpoint(c echo.Context) error {
	deck, e := db.Instance().Find(c.Param("id"))
	if e != nil {
		return helpers.NewHTTPError(
			http.StatusNotFound,
			"Invalid deck_id",
			"Could not find a deck with the desired ID",
		)
	}

	return c.JSON(http.StatusOK, deck)
}

func FetchDeckCardsEndpoint(c echo.Context) error {
	deck, e := db.Instance().Find(c.QueryParam("id"))
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

    cards := deck.Draw(int(amount))

    if len(cards.Cards) == 0 {
		return helpers.NewHTTPError(
			http.StatusNotFound,
            "Could not get more cards",
			"The deck is empty",
		)
    }

	return c.JSON(
		http.StatusOK,
        cards,
	)
}
