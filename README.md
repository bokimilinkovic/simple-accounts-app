Simple app that has main server that receives HTTP request for all CRUD operations for manipulation with accounts.

## e.GET("/accounts", accountHandler.GetAccounts) Returns all accounts with balances

curl -H 'Content-type: application/json' localhost:8080/accounts

## e.POST("/accounts", accountHandler.CreateAccount)

curl -X POST -H 'Content-type: application/json' -d '{"owner":"boki","balance":100}' localhost:8080/accounts

## e.GET("/account/:id", accountHandler.GetAccount)

curl -H 'Content-type: application/json' localhost:8080/account/{ownerName}

## e.DELETE("/accounts/:id", accountHandler.DeleteAccount)

curl -X DELETE -H 'Content-type: application/json' localhost:8080/account/{ownerName}

## e.PUT("/accounts/:id", accountHandler.UpdateAccount)

---- Creation of new transaction -----
We will simulate sending money from one account to another.
**curl -X POST -H 'Content-type: application/json' -d '{"sender":"boki","amount":29,"receiver":"john"}' localhost:8080/transactions**
Amount of money should be subtracted from sender's balances, and added to receiver. Only condition is that sender has enough of money on account.
At the same time, new transaction record is created and pushed to Redis MSQ.
Another application **transactions** waits for new messages on redis msq channel, once the transaction occurs, it's record is saved to separated database schema. In that way we created asychronous communication between (micro)services, so the latency is reduced...
**_What is left to do: Once the transaction is saved in DB and received ID, send it back to 'accounts' service, using another channel, or the same one. _**

## Create new database migration

migrate create -ext sql -dir db/migration -seq init_schema

## Run all services using docker-compose

docker-compose up --build
