# PlayingCards

A simple REST API to emulate the creation an manipulation of decks of cards.

## Project Structure

This project is structured based on the MVC architectural pattern. All files that are part of the project can be found under the 'src' folder, both the business logic as well as unit tests and a test for the endpoints can be found under this dir.

### models folder

The models folder contains the following files: suits.go, card.go, cards.go and deck.go.
All of these files come with a data structure representing the model itself, as well as helper functions and member functions.

### routes folder

The routes folder contain the routing to our endpoint '/decks', this folder presents a single file, decks.go that is responsible to relate a controller method to its respective endpoint with the goal of registering this route on the server, echo.

### helpers folder

Global helpers are kept on the helpers folder on the base of our project. The helpers found on this
project are used as facilitators for the creation of errors that are handled by the server API.

### controllers folder

The controllers folder handles all of the validations specific to a particular route, being them:
parsing query parameters or address parameters, inserting/retrieving content from the database or
preparing a properly formatted response.

### db folder

Mocks a database using a slice, the slice is a singleton and implements methods such as: insert, find, wipe and peek

## Installing dependencies

Assuming go is already installed on your environment, invoke make vendor

```sh
$ make vendor
```

## Running tests

Unit tests and integration tests can be executed by invoking make tests

```sh
$ make tests
```

## Notes on business rules

For the Create Deck Route, the cards passed as a query parameter are not sorted,
int other words, they keep the order of the cards as sent through the request.

For the entire project, 10 of any suit is written as 'X'. e.g.: XS, 10 of spades

## Routes

[Create Deck](./endpoints/create_deck.md)
[Fetch One Deck](./endpoints/fetch_one_deck.md)
[Fetch All Decks](./endpoints/fetch_all_decks.md)
[Draw Cards From Deck](./endpoints/draw_card_from_deck.md)
