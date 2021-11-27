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

// NewService returns the default Service interface implementation.
func NewService(repository storage.Repository) Service {
	return Service{repository: repository}
}
