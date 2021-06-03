package controllers

import (
    "errors"
    "strings"
    "strconv"
    "net/http"
    "github.com/labstack/echo/v4"
    "caiomcg.com/playing_cards/src/models"
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
        return e
    }

    return c.JSON(http.StatusOK, deck)
}

func FetchDeckCardsEndpoint(c echo.Context) error {
    id := c.Param("id")
    count, _ := strconv.ParseInt(c.Param("count"), 10, 32)

    deck, e := findDeck(id)

    if e != nil {
        return e
    }

    return c.JSON(
        http.StatusOK, 
        deck.Cards[0:getAvailableRange(int(count), len(deck.Cards))],
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
