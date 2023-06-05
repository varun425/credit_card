# Creadit Cards - POC
This project is a Proof of Concept (POC) for handling credit cards using a Go language server and CouchDB as the database. Below are the steps to set up and run the project.

DB Used : couchDB <br>
Lang: Go lang v1.19  <br>
port: 8000 <br>
DB Access url and credentials: http://localhost:5984/_utils<br>
user:admin <br>
password:adminpw

## Dependencies 
```
1 Docker: Required to run CouchDB as a container.
2 Go environment: Make sure you have Go installed on your machin
```
## Setup
```
1 Run the dbServer.sh script to set up the CouchDB server as a Docker container. This script will download and run the CouchDB image with the necessary configurations.
2 Execute the command go run main.go to start the Go server. This will run the server on port 8000.
```
## Postman Collection

In the same directory as the project, you can find a Postman collection JSON file named card_collection.json. 
Please make sure to import the collection into Postman and use it to interact with the server.

```
For reference below are the api endpoints 

Endpoint for submitting cards for admin: POST localhost:8000/admin/submitcards

Use this endpoint to submit a batch of credit cards for processing by the admin. The request body should contain an array of credit card objects.


Endpoint for submitting a card for user: POST localhost:8000/user/submitcard

Use this endpoint to submit a single credit card for processing by the user. The request body should contain the credit card details.


Endpoint to get information about a single card: GET localhost:8000/getsinglecard/{cardID}

Use this endpoint to retrieve information about a specific credit card. Replace {cardID} with the actual ID of the card.


Endpoint to get all the cards: GET localhost:8000/getallcards

Use this endpoint to retrieve information about all the credit cards stored in the database.

```