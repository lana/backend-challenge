package memory

import (
	"context"
	"patriciabonaldy/lana/internal/models"
	"patriciabonaldy/lana/internal/platform/storage"
	"sync"
)

// Memory is a memory Repository implementation.
type Memory struct {
	mux        sync.Mutex
	basketStge map[string]models.Basket
}

// NewRepository initializes a memory implementation of storage.Repository.
func NewRepository() storage.Repository {
	return &Memory{basketStge: make(map[string]models.Basket)}
}

// GetItem implements the storage.Repository interface.
func (m *Memory) GetItem(ctx context.Context, basketID string, productCode string) (models.Item, error) {
	defer m.mux.Unlock()

	m.mux.Lock()
	item, ok := m.basketStge[basketID].Items[productCode]
	if !ok {
		return models.Item{}, models.ErrItemNotFound
	}

	return item, nil
}

// UpdateBasket implements the storage.Repository interface.
func (m *Memory) UpdateBasket(ctx context.Context, basket models.Basket) (models.Basket, error) {
	defer m.mux.Unlock()

	m.mux.Lock()
	_, ok := m.basketStge[basket.Code]
	if !ok {
		return models.Basket{}, models.ErrBasketNotFound
	}

	m.basketStge[basket.Code] = basket

	return basket, nil
}

// CreateBasket implements the storage.Repository interface.
func (m *Memory) CreateBasket(ctx context.Context, id string) (models.Basket, error) {
	defer m.mux.Unlock()

	m.mux.Lock()
	if _, exist := m.basketStge[id]; exist {
		return models.Basket{}, models.ErrBasketCreated
	}

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
		return models.Basket{}, models.ErrBasketNotFound
	}

	return basket, nil
}

// RemoveBasket implements the storage.Repository interface.
func (m *Memory) RemoveBasket(ctx context.Context, basketID string) error {
	defer m.mux.Unlock()

	m.mux.Lock()
	_, ok := m.basketStge[basketID]
	if !ok {
		return models.ErrBasketNotFound
	}

	delete(m.basketStge, basketID)
	return nil
}

// RemoveProduct implements the storage.Repository interface.
func (m *Memory) RemoveProduct(ctx context.Context, basketID string, productCode string) (models.Basket, error) {
	defer m.mux.Unlock()

	m.mux.Lock()
	basket, ok := m.basketStge[basketID]
	if !ok {
		return models.Basket{}, models.ErrBasketNotFound
	}

	item, ok := basket.Items[productCode]
	if !ok {
		return models.Basket{}, models.ErrItemNotFound
	}

	delete(basket.Items, productCode)

	basket.Total -= item.Total
	m.basketStge[basketID] = basket

	return basket, nil
}
