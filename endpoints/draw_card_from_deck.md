## **Draw Card From Deck**

Returns json data about upon the creation of a deck.

- **URL**

  /decks/cards

- **Method:**

  `PUT`

- **URL Params**

  None

- **Query Params**

  required:

  `id=[uuid]` deck ID

  `amount=[int]` amount of cards to draw from the deck (at least one)

- **Data Params**

  None

- **Success Response:**

  - **Code:** 200 <br />
    **Content:** `{ "deck_id": "e5db62c5-fd0c-4cc9-a460-89030a1d836e", "shuffled": false, "remaining": 52, "cards": [ { "value": "ACE", "suit": "SPADES", "code": "AS" }, ...] }`

- **Fail Response**

  - **Code:** 404 <br />
    **Content:** `{ "error": "Invalid id", "message": "Could not find a deck with the desired ID" }`

  - **Code:** 404 <br />
    **Content:** `{ "error": "Could not get more cards", "message": "The deck is empty" }`

  - **Code:** 400 <br />
    **Content:** `{ "error": "Invalid amount", "message": "Amount not given or zero" }`
