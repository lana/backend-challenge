package memory

import (
	"context"
	"errors"
	"patriciabonaldy/lana/internal/models"
	"sync"
)

const (
	Pen    = "Pen"
	Tshirt = "Tshirt"
	Mug    = "Mug"
)

var (
	productMap = map[string]models.Product{
		Pen:    {Code: Pen, Name: "Lana Pen", Price: 5.00},
		Tshirt: {Code: Tshirt, Name: "Lana T-Shirt", Price: 20.00},
		Mug:    {Code: Mug, Name: "Lana Coffee Mug ", Price: 7.50},
	}

	ErrBasketCreated   = errors.New("basket was created previously")
	ErrBasketNotFound  = errors.New("basket does not exist")
	ErrProductNotFound = errors.New("product does not exist")
	ErrItemNotFound    = errors.New("item does not exist")
)

// Memory is a memory Repository implementation.
type Memory struct {
	mux        sync.Mutex
	basketStge map[string]models.Basket
}

// NewRepository initializes a memory implementation of storage.Repository.
func NewRepository() *Memory {
	return &Memory{basketStge: make(map[string]models.Basket)}
}

// AddProduct implements the storage.Repository interface.
func (m *Memory) AddProduct(ctx context.Context, basketID string, productCode string) (models.Basket, error) {
	defer m.mux.Unlock()

	m.mux.Lock()
	basket, ok := m.basketStge[basketID]
	if !ok {
		return models.Basket{}, ErrBasketNotFound
	}

	product, ok := productMap[productCode]
	if !ok {
		return models.Basket{}, ErrProductNotFound
	}

	item, ok := basket.Items[product.Code]
	if !ok {
		basket.Items[productCode] = models.Item{
			Product:  product,
			Quantity: 1,
		}

		return basket, nil
	}

	item.Quantity++
	basket.Items[productCode] = item

	return basket, nil
}

// CreateBasket implements the storage.Repository interface.
func (m *Memory) CreateBasket(ctx context.Context, id string) (models.Basket, error) {
	defer m.mux.Unlock()

	m.mux.Lock()
	/*id, err := uuid.NewUUID()
	if err != nil {
		return models.Basket{}, err
	}
	*/
	basket := models.NewBasket(id)
	m.basketStge[basket.Code] = basket
	return basket, nil
}

// FindBasketByID implements the storage.Repository interface.
func (m *Memory) FindBasketByID(ctx context.Context, id string) (models.Basket, error) {
	defer m.mux.Unlock()

	m.mux.Lock()
	basket, ok := m.basketStge[id]
	if !ok {
		return models.Basket{}, ErrBasketNotFound
	}

	return basket, nil
}

// RemoveProduct implements the storage.Repository interface.
func (m *Memory) RemoveProduct(ctx context.Context, basketID string, productCode string) (models.Basket, error) {
	defer m.mux.Unlock()

	m.mux.Lock()
	basket, ok := m.basketStge[basketID]
	if !ok {
		return models.Basket{}, ErrBasketNotFound
	}

	product, ok := productMap[productCode]
	if !ok {
		return models.Basket{}, ErrProductNotFound
	}

	_, ok = basket.Items[product.Code]
	if !ok {
		return models.Basket{}, ErrItemNotFound
	}

	delete(basket.Items, product.Code)

	return basket, nil
}
