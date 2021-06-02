package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

import "github.com/gorilla/mux"
var Decks []Deck

func returnAllDecks(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint hit: returnAllDecks")
    json.NewEncoder(w).Encode(Decks)
}

func handleRequest() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/decks", returnAllDecks)

    log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
    Decks = []Deck {
        CreateDeck(false),
        CreateDeck(true),
    }

    handleRequest()
}
