package memory_test

import (
	"context"
	"errors"
	"patriciabonaldy/lana/internal/models"
	"patriciabonaldy/lana/internal/platform/storage/platform/storage/memory"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemory_AddProduct_and_basket_does_not_exit(t *testing.T) {
	repository := memory.NewRepository()
	id, _ := uuid.NewUUID()
	_, err := repository.AddProduct(context.Background(), id.String(), "PEN")
	assert.Equal(t, errors.Is(err, models.ErrBasketNotFound), true)
}

func TestMemory_AddProduct(t *testing.T) {
	testcases := []struct {
		name          string
		item          string
		expectedError error
		want          models.Basket
	}{
		{
			name: "insert_tshirt",
			item: "Tshirt",
			want: models.Basket{
				Code: "",
				Items: map[string]models.Item{
					"Tshirt": {
						Product: models.Product{
							Code:  "Tshirt",
							Name:  "Lana T-Shirt",
							Price: 20,
						},
						Quantity: 1,
						Discount: 0,
					},
				},
				Total: 0,
			},
		},
		{
			name: "insert__another_tshirt",
			item: "Tshirt",
			want: models.Basket{
				Code: "",
				Items: map[string]models.Item{
					"Tshirt": {
						Product: models.Product{
							Code:  "Tshirt",
							Name:  "Lana T-Shirt",
							Price: 20,
						},
						Quantity: 2,
						Discount: 0,
					},
				},
				Total: 0,
			},
		},
		{
			name:          "does_not_exist_product",
			item:          "DRESS",
			expectedError: errors.New("product does not exist"),
			want:          models.Basket{},
		},
	}
	repository := memory.NewRepository()
	id, err := uuid.NewUUID()
	require.NoError(t, err)
	basket, err := repository.CreateBasket(context.Background(), id.String())
	require.NoError(t, err)
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			product, err := repository.AddProduct(context.Background(), basket.Code, test.item)
			if test.expectedError == nil {
				test.want.Code = id.String()
			}
			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.want, product)
		})
	}
}

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

func TestMemory_RemoveBaskett(t *testing.T) {
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
	var _, err = repository.CreateBasket(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256")

	_, err = repository.AddProduct(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256", "Tshirt")
	require.NoError(t, err)

	expectedError := errors.New("basket does not exist")
	_, err = repository.RemoveProduct(context.Background(), "999999999999-11ec-a386-1e003b1e5256", "Tshirt")
	assert.Equal(t, expectedError, err)

	expectedError = errors.New("product does not exist")
	_, err = repository.RemoveProduct(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256", "tshirt")
	assert.Equal(t, expectedError, err)

	want := models.Basket{
		Code:  "4200f350-4fa5-11ec-a386-1e003b1e5256",
		Items: make(map[string]models.Item),
		Total: 0,
	}
	basket, err := repository.RemoveProduct(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256", "Tshirt")
	assert.NoError(t, err)
	assert.Equal(t, want, basket)

	expectedError = errors.New("item does not exist")
	_, err = repository.RemoveProduct(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256", "Tshirt")
	assert.Equal(t, expectedError, err)
}
