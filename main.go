package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chibuikekenneth/go-transactions/app"
	"github.com/chibuikekenneth/go-transactions/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/transactions/new", controllers.CreateTransaction).Methods("POST")
	router.HandleFunc("/api/me/transactions", controllers.GetTransactionsFor).Methods("GET") //  user/2/contacts

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
