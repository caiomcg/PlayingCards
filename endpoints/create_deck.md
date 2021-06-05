**Create Deck**
----
  Returns json data about upon the creation of a deck.

* **URL**

  /decks

* **Method:**

  `POST`

*  **URL Params**

   None
  
*  **Query Params**

   `shuffle=[bool]` wether to shuffle the deck or not (default: false)
   `cards=[strings]` card codes to be added, as is, to the deck (default: all)

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ id : 12, name : "Michael Bloom" }`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "User doesn't exist" }`

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`

* **Sample Call:**

  ```javascript
    $.ajax({
      url: "/users/1",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ```
