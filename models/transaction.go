package models

import (
	"fmt"

	u "github.com/chibuikekenneth/go-transactions/utils"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	Name   string `json:"name"`
	Amount string `json:"amount"`
	UserId uint   `json:"user_id"` //The user that this amount belongs to
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (transaction *Transaction) Validate() (map[string]interface{}, bool) {

	if transaction.Name == "" {
		return u.Message(false, "Transaction name should be on the payload"), false
	}

	if transaction.Amount == "" {
		return u.Message(false, "Amount should be on the payload"), false
	}

	if transaction.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (transaction *Transaction) Create() map[string]interface{} {

	if resp, ok := transaction.Validate(); !ok {
		return resp
	}

	GetDB().Create(transaction)

	resp := u.Message(true, "success")
	resp["transaction"] = transaction
	return resp
}

func GetTransaction(id uint) *Transaction {

	transaction := &Transaction{}
	err := GetDB().Table("transactions").Where("id = ?", id).First(transaction).Error
	if err != nil {
		return nil
	}
	return transaction
}

func GetTransactions(user uint) []*Transaction {

	transactions := make([]*Transaction, 0)
	err := GetDB().Table("transactions").Where("user_id = ?", user).Find(&transactions).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return transactions
}
