## **Fetch Deck**

Returns a single Deck of cards based on the given uuid

- **URL**

  /decks/:id

- **Method:**

  `GET`

- **URL Params**

  None

- **Query Params**

  None

- **Data Params**

  None

- **Success Response:**

  - **Code:** 200 <br />
    **Content:** `{ "deck_id": "e5db62c5-fd0c-4cc9-a460-89030a1d836e", "shuffled": false, "remaining": 52, "cards": [ { "value": "ACE", "suit": "SPADES", "code": "AS" }, ...] }`

- **Fail Response**
  - **Code:** 404 <br />
    **Content:** `{ "error": "Invalid deck_id", "message": "Could not find a deck with the desired ID" }`
