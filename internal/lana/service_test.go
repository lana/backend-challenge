package lana

import (
	"context"
	"patriciabonaldy/lana/internal/models"
	"patriciabonaldy/lana/internal/platform/storage/platform/storage/storagemocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_GetBasketBasket_RepositoryError(t *testing.T) {
	repositoryMock := new(storagemocks.Repository)
	repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(models.Basket{}, models.ErrBasketNotFound)

	service := NewService(repositoryMock)
	_, err := service.GetBasket(context.Background(), "99999")

	repositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func TestService_GetBasketBasket_Success(t *testing.T) {
	repositoryMock := new(storagemocks.Repository)
	basketExpected := models.Basket{
		Code:  "99999",
		Total: 0,
	}
	repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(basketExpected, nil)

	service := NewService(repositoryMock)
	basket, err := service.GetBasket(context.Background(), "99999")

	repositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, basketExpected, basket)
}
