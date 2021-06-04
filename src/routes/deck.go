package routes

import "net/http"
import "fmt"
import "github.com/labstack/echo/v4"
import "caiomcg.com/playing_cards/src/controllers"

var baseRoute = "/decks"

func RegisterDeck(r *echo.Router) {
	r.Add(http.MethodPost, baseRoute, controllers.CreateDeckEndpoint)
	r.Add(http.MethodGet, baseRoute, controllers.FetchDecksEndpoint)
	r.Add(
		http.MethodGet,
		fmt.Sprintf("%s/:id", baseRoute),
		controllers.OpenDeckEndpoint,
	)
	r.Add(
		http.MethodGet,
		fmt.Sprintf("%s/cards", baseRoute),
		controllers.FetchDeckCardsEndpoint,
	)
}
