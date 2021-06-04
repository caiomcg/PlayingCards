package main

import (
	"caiomcg.com/playing_cards/src/helpers"
	"caiomcg.com/playing_cards/src/routes"
	"fmt"
	"log"
	"os"
)

import "github.com/joho/godotenv"
import "github.com/labstack/echo/v4"
import "github.com/labstack/echo/v4/middleware"

func CreateServer() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = helpers.ErrorHandler

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))

	e.Use(middleware.Recover())
	routes.RegisterDeck(e.Router())

	return e
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro loading .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	e := CreateServer()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
