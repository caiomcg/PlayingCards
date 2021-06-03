package main

import (
    "caiomcg.com/playing_cards/src/routes"
    "caiomcg.com/playing_cards/src/helpers"
)

import "github.com/labstack/echo/v4"
import "github.com/labstack/echo/v4/middleware"

func createAndStartServer() {
    e := echo.New()
    e.HTTPErrorHandler = helpers.ErrorHandler

    e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
      Format: "method=${method}, uri=${uri}, status=${status}\n",
    }))
    e.Use(middleware.Recover())

    routes.RegisterDeck(e.Router())

    e.Logger.Fatal(e.Start(":8000"))
}

func main() {
    createAndStartServer()
}
