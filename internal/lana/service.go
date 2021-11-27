package lana

import (
	"context"
	"errors"
	"patriciabonaldy/lana/internal/models"
	"patriciabonaldy/lana/internal/platform/storage"

	"github.com/google/uuid"
)

var ErrIDIsRequired = errors.New("basket id is required")

// Service is the default lanaService interface
// implementation returned by lana.NewService.
type Service struct {
	repository storage.Repository
}

type ProductRequest struct {
	BasketID    string `json:"basket_id" binding:"required"`
	ProductCode string `json:"product_code" binding:"required"`
}

// NewService returns the default Service interface implementation.
func NewService(repository storage.Repository) Service {
	return Service{repository: repository}
}

func (s Service) CreateBasket(ctx context.Context) (models.Basket, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return models.Basket{}, err
	}

	basket, err := s.repository.CreateBasket(ctx, id.String())
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}

func (s Service) GetBasket(ctx context.Context, id string) (models.Basket, error) {
	basket, err := s.repository.FindBasketByID(ctx, id)
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}

func (s Service) RemoveBasket(ctx context.Context, id string) error {
	err := s.repository.RemoveBasket(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) AddProduct(ctx context.Context, request ProductRequest) (models.Basket, error) {
	basket, err := s.repository.AddProduct(ctx, request.BasketID, request.ProductCode)
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}

func (s Service) RemoveProduct(ctx context.Context, request ProductRequest) (models.Basket, error) {
	var basket models.Basket

	basket, err := s.repository.RemoveProduct(ctx, request.BasketID, request.ProductCode)
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}
