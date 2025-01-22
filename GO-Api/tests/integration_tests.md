# Integration Tests for Library/Bookstore API
Initiate GO server

go run main.go


## 1. Get All Books
```bash

curl localhost:8080/books



## 2. Get Book ID 2
curl localhost:8080/books/2



# 3. Checkout Book with ID 2
curl "localhost:8080/checkout?id=2" --request "PATCH"



# 4. Return Book with ID 2
curl "localhost:8080/return?id=2" --request "PATCH"


# 5. create a new book with ID 7
curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"



# 6. checkout book with no ID provided
curl "http://localhost:8080/checkout" --request "PATCH"