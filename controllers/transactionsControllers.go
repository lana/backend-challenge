package controllers

import (
	"encoding/json"
	"net/http"

	u "github.com/chibuikekenneth/go-transactions/utils"

	"github.com/chibuikekenneth/go-transactions/models"
)

var CreateTransaction = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that sends the request
	transaction := &models.Transaction{}

	err := json.NewDecoder(r.Body).Decode(transaction)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	transaction.UserId = user
	resp := transaction.Create()
	u.Respond(w, resp)
}

var GetTransactionsFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetTransactions(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
