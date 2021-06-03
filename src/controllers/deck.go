package controllers

import (
    "errors"
    "strings"
    "strconv"
    "net/http"
    "github.com/labstack/echo/v4"
    "caiomcg.com/playing_cards/src/models"
    "caiomcg.com/playing_cards/src/helpers"
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

    if (deckDTO.ID == "" || deckDTO.Amount == 0) {
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

// Helpers
func findDeck(id string) (models.Deck, error) {
    for _, v := range Decks {
        if v.Id.String() == id {
            return v, nil
        }
    }
    return models.Deck{}, errors.New("Deck not found")
}

func processShuffleParam(shuffle string) bool {
    if (shuffle == "") {
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
