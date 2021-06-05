# PlayingCards

This project is part of the Unattended Programming Test from Toggl with the requirements to create
a backend REST API that handles a Deck of cards. This project has 3 endpoints that can be used
to create and fetch Decks or Cards from a Deck.

## Project Structure

This project is divide in a MVC like project structure where, under the src folder can be found
both the business logic of the project, as well as the tests.

### models folder

The models folder contains the following files: suits.go, card.go, cards.go and deck.go.
All of these files come with a data structure representing the model of each entity and some
business rules that can be applied on this model or to help create and manipulate such model.

### routes folder

The routes folder contain the routing to our endpoint '/decks', this folder presents a single file,
decks.go that is responsible to relate a controller method to its respective endpoint with the goal
of registering this route on the server, echo.

### helpers folder

Global helpers are kept on the helpers folder on the base of our project. The helpers found on this
project are use as a facilitator to the creation of errors that are handled by the server API.

### controllers folder

The controllers folder handles all of the validations specific to a particular route, being them:
parsing query parameters or address parameters, inserting/retrieving content from the database or
preparing a properly formatted response

## Notes

To make the process of writing this particular project faster I have opted to not use a database,
being it a relational/non relational one or an in memory database. The 'database' used on this
project is just an array that is looped over to access content or inserted if a POST is requested.
I am aware that this do not provide ACID support and am assuming this will not be a problem that
is evaluated on this assignment.

## Routes
