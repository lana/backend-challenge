# GO Transactions

GO Transaction is an API which stores a complete transaction history for transactions between users.

This Application uses PostgreSQL as its Primary Database.

# Dependencies
[gorilla/mux ](https://github.com/gorilla/mux)— A powerful URL router and dispatcher. 
[jinzhu/gorm ](https://github.com/jinzhu/gorm)— The fantastic ORM library for Golang
[dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go) — Used to sign and verify JWT tokens
[joho/godotenv](https://github.com/joho/godotenv) — Used to load .env files into the project

To install any of this package, open terminal and run

go get github.com/{package-name}

### Authentication

With JWT, we created a unique token for each authenticated user, this token was included in the header of the subsequent request made to the API server, this method allow us to identify every users that make calls to our API. 

## Deployment




