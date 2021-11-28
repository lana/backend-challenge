package memory_test

import (
	"context"
	"errors"
	"patriciabonaldy/lana/internal/models"
	"patriciabonaldy/lana/internal/platform/storage/platform/storage/memory"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemory_FindBasketByID(t *testing.T) {
	testcases := []struct {
		name          string
		basketID      string
		expectedError error
		want          models.Basket
	}{
		{
			name:          "Found_ok",
			basketID:      "4200f350-4fa5-11ec-a386-1e003b1e5256",
			expectedError: nil,
			want: models.Basket{
				Code:  "4200f350-4fa5-11ec-a386-1e003b1e5256",
				Items: make(map[string]models.Item),
				Total: 0,
			},
		},
		{
			name:          "Not_found",
			basketID:      "99999999999999-a386-1e003b1e5256",
			expectedError: errors.New("basket does not exist"),
			want: models.Basket{
				Code:  "",
				Total: 0,
			},
		},
	}

	repository := memory.NewRepository()
	_, err := repository.CreateBasket(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256")
	require.NoError(t, err)
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			basket, err := repository.FindBasketByID(context.Background(), test.basketID)
			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.want, basket)
		})
	}
}

func TestMemory_GetItem(t *testing.T) {
	repository := memory.NewRepository()
	ctx := context.Background()

	basket, err := repository.CreateBasket(ctx, "4200f350-4fa5-11ec-a386-1e003b1e5256")
	require.NoError(t, err)

	_, err = repository.GetItem(ctx, basket.Code, "PEN")
	assert.EqualError(t, err, "item does not exist")

	basket.Items = map[string]models.Item{
		"Tshirt": {
			Product: models.Product{
				Code:  "Tshirt",
				Name:  "Lana T-Shirt",
				Price: 20,
			},
			Quantity: 1,
			Total:    20,
		},
	}
	basket.Total = 20
	_, err = repository.UpdateBasket(ctx, basket)
	require.NoError(t, err)

	_, err = repository.GetItem(ctx, basket.Code, "Tshirt")
	assert.NoError(t, err)
}

func TestMemory_RemoveBasket(t *testing.T) {
	repository := memory.NewRepository()
	_, err := repository.CreateBasket(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256")
	require.NoError(t, err)

	expectedError := errors.New("basket does not exist")
	err = repository.RemoveBasket(context.Background(), "999999999999-11ec-a386-1e003b1e5256")
	assert.Equal(t, expectedError, err)

	err = repository.RemoveBasket(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256")
	assert.NoError(t, err)
}

func TestMemory_RemoveProduct(t *testing.T) {
	repository := memory.NewRepository()
	ctx := context.Background()

	expectedError := errors.New("basket does not exist")
	_, err := repository.RemoveProduct(ctx, "999999999999-11ec-a386-1e003b1e5256", "Tshirt")
	assert.Equal(t, expectedError, err)

	basket, err := repository.CreateBasket(ctx, "4200f350-4fa5-11ec-a386-1e003b1e5256")
	require.NoError(t, err)

	expectedError = errors.New("item does not exist")
	_, err = repository.RemoveProduct(ctx, "4200f350-4fa5-11ec-a386-1e003b1e5256", "tshirt")
	assert.Equal(t, expectedError, err)

	expectedError = errors.New("item does not exist")
	_, err = repository.RemoveProduct(ctx, "4200f350-4fa5-11ec-a386-1e003b1e5256", "Tshirt")
	assert.Equal(t, expectedError, err)

	basket.Items = map[string]models.Item{
		"Tshirt": {
			Product: models.Product{
				Code:  "Tshirt",
				Name:  "Lana T-Shirt",
				Price: 20,
			},
			Quantity: 1,
			Total:    20,
		},
	}
	basket.Total = 20
	_, err = repository.UpdateBasket(ctx, basket)
	require.NoError(t, err)

	_, err = repository.RemoveProduct(ctx, "4200f350-4fa5-11ec-a386-1e003b1e5256", "Tshirt")
	assert.NoError(t, err)
}
