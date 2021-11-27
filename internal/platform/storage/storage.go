package storage

import (
	"context"
	"patriciabonaldy/lana/internal/models"
)

// Repository defines the expected behaviour from a lana storage.
type Repository interface {
	FindBasketByID(ctx context.Context, id string) (models.Basket, error)
	CreateBasket(ctx context.Context, id string) (models.Basket, error)
	AddProduct(ctx context.Context, basketID, productCode string) (models.Basket, error)
	RemoveProduct(ctx context.Context, basketID, productCode string) (models.Basket, error)
	RemoveBasket(ctx context.Context, id string) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=Repository
