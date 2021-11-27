package models

type Basket struct {
	Code  string
	Items map[string]Item
	Total float64
}

type Product struct {
	Code  string
	Name  string
	Price float64
}

type Item struct {
	Product  Product
	Quantity int
	Discount float64
}

func NewBasket(id string) Basket {
	return Basket{
		Code:  id,
		Items: make(map[string]Item),
	}
}
