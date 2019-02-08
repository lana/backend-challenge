# GO Transactions

GO Transaction is an API which stores a complete transaction history for transact  ions between users.

This Application uses PostgreSQL as its Primary Database.

# Dependencies
1. gorilla/mux — A powerful URL router and dispatcher. 
2. jinzhu/gorm — The fantastic ORM library for Golang
3. dgrijalva/jwt-go — Used to sign and verify JWT tokens
4. joho/godotenv — Used to load .env files into the project

To install any of this package, open terminal and run

go get github.com/{package-name}

### Authentication

With JWT, we created a unique token for each authenticated user, this token was included in the header of the subsequent request made to the API server, this method allow us to identify every users that make calls to our API. 

## Deployment




