package main

import "testing"
import (
	"caiomcg.com/playing_cards/src/controllers"
	"caiomcg.com/playing_cards/src/models"
	"encoding/json"
	"fmt"
	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
)

var (
	validDeck           = `{}`
	STANDARD_SIZE uint8 = 52
)

func resetDb() {
	controllers.Decks = []models.Deck{
		models.CreateDeck(true, []string{}),
		models.CreateDeck(false, []string{"AS", "2C"}),
	}
}

func extractDeck(data []byte) models.Deck {
	var deck models.Deck

	if err := json.Unmarshal(data, &deck); err != nil {
		panic(err)
	}

	return deck
}

func extractCards(data []byte) models.Cards {
	var cards models.Cards

	if err := json.Unmarshal(data, &cards); err != nil {
		panic(err)
	}

	return cards
}

func TestCreateDeckEndpoint(t *testing.T) {
	r := gofight.New()

	// Regular request, unshuffled deck with 52 cards
	r.POST("/decks").
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)

			deck := extractDeck(r.Body.Bytes())

			assert.Equal(t, STANDARD_SIZE, uint8(len(deck.Cards)))
			assert.Equal(t, false, deck.Shuffled)
			assert.Equal(t, STANDARD_SIZE, deck.Remaining)
		})

	// Explicitly Querying for an unshuffled deck
	r.POST("/decks").
		SetQuery(gofight.H{
			"shuffle": "false",
		}).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)

			data := []byte(r.Body.String())
			var deck models.Deck
			if err := json.Unmarshal(data, &deck); err != nil {
				panic(err)
			}

			assert.Equal(t, STANDARD_SIZE, uint8(len(deck.Cards)))
			assert.Equal(t, false, deck.Shuffled)
			assert.Equal(t, STANDARD_SIZE, deck.Remaining)
		})

	// Querying for a shuffled deck
	r.POST("/decks").
		SetQuery(gofight.H{
			"shuffle": "true",
		}).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)

			data := []byte(r.Body.String())
			var deck models.Deck
			if err := json.Unmarshal(data, &deck); err != nil {
				panic(err)
			}

			assert.Equal(t, STANDARD_SIZE, uint8(len(deck.Cards)))
			assert.Equal(t, true, deck.Shuffled)
			assert.Equal(t, STANDARD_SIZE, deck.Remaining)
		})

	// Querying for an unshuffled deck with specific cards
	r.POST("/decks").
		SetQuery(gofight.H{
			"cards": "AS,QH,KD",
		}).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)

			data := []byte(r.Body.String())
			var deck models.Deck
			if err := json.Unmarshal(data, &deck); err != nil {
				panic(err)
			}

			assert.Equal(t, 3, len(deck.Cards))
			assert.Equal(t, false, deck.Shuffled)
			assert.Equal(t, uint8(3), deck.Remaining)
		})

	// Querying for a shuffled deck with specific cards
	r.POST("/decks").
		SetQuery(gofight.H{
			"shuffle": "true",
			"cards":   "AS,2S,3S,4S,5S,6S,7S,8S,9S,XS",
		}).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)

			data := []byte(r.Body.String())
			var deck models.Deck
			if err := json.Unmarshal(data, &deck); err != nil {
				panic(err)
			}

			assert.Equal(t, 10, len(deck.Cards))
			assert.Equal(t, true, deck.Shuffled)
			assert.Equal(t, uint8(10), deck.Remaining)
		})

	// Querying for a shuffled deck with duplicate cards
	r.POST("/decks").
		SetQuery(gofight.H{
			"shuffle": "true",
			"cards":   "AS,2S,AS",
		}).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)

			data := []byte(r.Body.String())
			var deck models.Deck
			if err := json.Unmarshal(data, &deck); err != nil {
				panic(err)
			}

			assert.Equal(t, 2, len(deck.Cards))
			assert.Equal(t, true, deck.Shuffled)
			assert.Equal(t, uint8(2), deck.Remaining)
		})

	// Querying for a shuffled deck with more than 52 cards
	r.POST("/decks").
		SetQuery(gofight.H{
			"shuffle": "true",
			"cards":   `AS,2S,3S,4S,5S,6S,7S,8S,9S,XS,JS,QS,KS,AC,2C,3C,4C,5C,6C,7C,8C,9C,XC,JC,QC,KC,AD,2D,3D,4D,5D,6D,7D,8D,9D,XD,JD,QD,KD,AH,2H,3H,4H,5H,6H,7H,8H,9H,XH,JH,QH,KH,AH,2H,3H,4H,5H,6H,7H,8H,9H,XH,JH,QH,KH`,
		}).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)

			data := []byte(r.Body.String())
			var deck models.Deck
			if err := json.Unmarshal(data, &deck); err != nil {
				panic(err)
			}

			assert.Equal(t, int(STANDARD_SIZE), len(deck.Cards))
			assert.Equal(t, true, deck.Shuffled)
			assert.Equal(t, STANDARD_SIZE, deck.Remaining)
		})
}

func TestFetchDecksEndpoint(t *testing.T) {
	r := gofight.New()

	resetDb()

	r.GET("/decks").
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestOpenDeckEndpoint(t *testing.T) {
	r := gofight.New()

	resetDb()

	r.GET(fmt.Sprintf("/decks/%s", controllers.Decks[0].Id)).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)

			data := []byte(r.Body.String())
			var deck models.Deck
			if err := json.Unmarshal(data, &deck); err != nil {
				panic(err)
			}

			assert.Equal(t, controllers.Decks[0].Id, deck.Id)
		})

	r.GET(fmt.Sprintf("/decks/%s", "INVALID")).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusNotFound, r.Code)
		})
}

func TestFetchDeckCardsEndpoint(t *testing.T) {
	r := gofight.New()

	resetDb()

	// Request with missing amount
	r.GET("/decks/cards").
		SetQuery(gofight.H{
			"id":     controllers.Decks[0].Id.String(),
			"amount": "10",
		}).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)

			cards := extractCards(r.Body.Bytes())

			assert.Equal(t, 10, len(cards.Cards))
		})

	// Request witouth a valid body
	r.GET("/decks/cards").
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusNotFound, r.Code)
		})

	// Request with missing keys
	r.GET("/decks/cards").
		SetQuery(gofight.H{
			"id":    "INVALID",
			"count": "10",
		}).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusNotFound, r.Code)
		})

	// Request with invalid deck id
	r.GET("/decks/cards").
		SetQuery(gofight.H{
			"id":     "INVALID",
			"amount": "2",
		}).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusNotFound, r.Code)
		})

	// Request with missing amount
	r.GET("/decks/cards").
		SetQuery(gofight.H{
			"id": controllers.Decks[0].Id.String(),
		}).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})

	// Request with invalid amount
	r.GET("/decks/cards").
		SetQuery(gofight.H{
			"id":     controllers.Decks[0].Id.String(),
			"amount": "0",
		}).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})

	r.GET("/decks/cards").
		SetQuery(gofight.H{
			"id":     controllers.Decks[0].Id.String(),
			"amount": "-12",
		}).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})

	r.GET("/decks/cards").
		SetQuery(gofight.H{
			"id":     controllers.Decks[0].Id.String(),
			"amount": "INVALID",
		}).
		Run(CreateServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
}
