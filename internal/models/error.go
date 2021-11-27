package models

import "errors"

var (
	ErrBasketCreated   = errors.New("basket was created previously")
	ErrBasketNotFound  = errors.New("basket does not exist")
	ErrProductNotFound = errors.New("product does not exist")
	ErrItemNotFound    = errors.New("item does not exist")
)
