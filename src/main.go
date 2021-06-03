package main

import (
    "caiomcg.com/playing_cards/src/routes"
)

import "github.com/labstack/echo/v4"
import "github.com/labstack/echo/v4/middleware"


func registerRoutes(r *echo.Router) {
    routes.RegisterDeck(r)
}

func createAndStartServer() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    routes.RegisterDeck(e.Router())

    e.Logger.Fatal(e.Start(":8000"))
}

func main() {
    createAndStartServer()
}
