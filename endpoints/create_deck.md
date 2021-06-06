## **Create Deck**

Returns json data upon the creation of a deck.

- **URL**

  /decks

- **Method:**

  `POST`

- **URL Params**

  None

- **Query Params**

  `shuffle=[bool]` wether to shuffle the deck or not (default: false)

  `cards=[strings]` card codes to be added, as is, to the deck (default: all)

- **Data Params**

  None

- **Success Response:**

  - **Code:** 200 <br />
    **Content:** `{ "deck_id": "e5db62c5-fd0c-4cc9-a460-89030a1d836e", "shuffled": false, "remaining": 52, "cards": [ { "value": "ACE", "suit": "SPADES", "code": "AS" }, ...] }`

- **Fail Response:**
  - **Code:** 400 <br />
    **Content:** `{ "error": "Invalid custom card", "message": "Card codes should follow the estipulated rules" }`
