package models

const (
	Pen                    = "Pen"
	Tshirt                 = "Tshirt"
	Mug                    = "Mug"
	PenQuantity            = 2
	QuantityDiscount       = 0.25
	TshirtQuantityDiscount = 3
)

var (
	ProductMap = map[string]Product{
		Pen:    {Code: Pen, Name: "Lana Pen", Price: 5.00},
		Tshirt: {Code: Tshirt, Name: "Lana T-Shirt", Price: 20.00},
		Mug:    {Code: Mug, Name: "Lana Coffee Mug ", Price: 7.50},
	}
)

type Basket struct {
	Code  string
	Items map[string]Item
	Total float64
	Close bool
}

type Product struct {
	Code  string
	Name  string
	Price float64
}

type Item struct {
	Product  Product
	Quantity int
	Total    float64
}

func NewBasket(id string) Basket {
	return Basket{
		Code:  id,
		Items: make(map[string]Item),
	}
}

func (b *Basket) CalculateTotal() {
	var total float64
	for _, i := range b.Items {
		total += i.Total
	}

	b.Total = total
}

func (i *Item) WithOutDiscount() {
	var discountAmount float64

	product := i.Product
	discountAmount = product.Price * float64(i.Quantity)
	i.Total = discountAmount
}

// Discount3OrMore function
// Check if client buy 3 or more the same type
// apply 25% discount over amount
func Discount3OrMore(item Item) Item {
	var discountAmount float64
	product := item.Product
	if item.Quantity >= TshirtQuantityDiscount {
		discountAmount = (product.Price * float64(item.Quantity)) - ((product.Price * float64(item.Quantity)) * QuantityDiscount)
		item.Total = discountAmount

		return item
	}

	discountAmount = product.Price * float64(item.Quantity)
	item.Total = discountAmount

	return item
}

// DiscountBuyingTwoGetOneFree function
// Check if client buy 2 or more the same type
// gift one free
func DiscountBuyingTwoGetOneFree(item Item) Item {
	product := item.Product
	total := product.Price * float64(item.Quantity)
	if item.Quantity >= PenQuantity {
		item.Quantity++
	}

	item.Total = total

	return item
}
