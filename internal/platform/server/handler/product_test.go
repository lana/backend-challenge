package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"patriciabonaldy/lana/internal/lana"
	"patriciabonaldy/lana/internal/models"
	"patriciabonaldy/lana/internal/platform/storage/platform/storage/storagemocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAddProductHandler(t *testing.T) {
	request := lana.ProductRequest{
		BasketID:    "4200f350-4fa5-11ec-a386-1e003b1e5256",
		ProductCode: "Tshirt",
	}

	gin.SetMode(gin.TestMode)

	t.Run("given a valid id request it returns 200", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		itemMock := models.Item{
			Product: models.Product{
				Code:  "Tshirt",
				Name:  "Lana T-Shirt",
				Price: 20,
			},
		}

		itemMock.Quantity = 1
		itemMock.Total = 20
		basketExpected := models.Basket{
			Code: "4200f350-4fa5-11ec-a386-1e003b1e5256",
			Items: map[string]models.Item{
				"Tshirt": itemMock,
			},
			Total: 20,
		}
		basketmock := models.Basket{
			Code:  "4200f350-4fa5-11ec-a386-1e003b1e5256",
			Items: make(map[string]models.Item),
			Total: 0,
		}
		repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(basketmock, nil).Once()
		repositoryMock.On("GetItem", mock.Anything, mock.Anything, mock.Anything).Return(models.Item{}, models.ErrItemNotFound).Once()
		repositoryMock.On("CreateItem", mock.Anything, mock.Anything, mock.Anything).Return(itemMock, nil)
		repositoryMock.On("UpdateBasket", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(basketExpected, nil).Once()
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.POST("/baskets/:id/products/:code", AddProductHandler(service))

		url := fmt.Sprintf("/baskets/%s/products/%s", request.BasketID, request.ProductCode)
		req, err := http.NewRequest(http.MethodPost, url, nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("given a invalid BasketID it returns 400", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything, mock.Anything).
			Return(models.Basket{}, models.ErrBasketNotFound)

		service := lana.NewService(repositoryMock)
		r := gin.New()
		r.POST("/baskets/:id/products/:code", AddProductHandler(service))

		url := fmt.Sprintf("/baskets/%s/products/%s", request.BasketID, request.ProductCode)
		req, err := http.NewRequest(http.MethodPost, url, nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

func TestRemoveProductHandler(t *testing.T) {
	request := lana.ProductRequest{
		BasketID:    "4200f350-4fa5-11ec-a386-1e003b1e5256",
		ProductCode: "Tshirt",
	}
	basketExpected := models.Basket{
		Code:  "4200f350-4fa5-11ec-a386-1e003b1e5256",
		Items: make(map[string]models.Item),
		Total: 0,
	}

	gin.SetMode(gin.TestMode)

	t.Run("given a valid request it returns 200", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(basketExpected, nil)
		repositoryMock.On("RemoveProduct", mock.Anything, mock.Anything, mock.Anything).Return(basketExpected, nil)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.DELETE("/baskets/:id/products/:code", RemoveProductHandler(service))

		url := fmt.Sprintf("/baskets/%s/products/%s", request.BasketID, request.ProductCode)
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("given a invalid product inside request it returns 400", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(basketExpected, nil)
		repositoryMock.On("RemoveProduct", mock.Anything, mock.Anything, mock.Anything).Return(basketExpected, models.ErrProductNotFound)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.DELETE("/baskets/:id/products/:code", RemoveProductHandler(service))

		url := fmt.Sprintf("/baskets/%s/products/%s", request.BasketID, request.ProductCode)
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
	t.Run("given a invalid request it returns 400", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("RemoveProduct", mock.Anything, mock.Anything, mock.Anything).Return(basketExpected, models.ErrProductNotFound)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.DELETE("/baskets/:id/products/:code", RemoveProductHandler(service))

		url := fmt.Sprintf("/baskets/%s/products/%s", "6767868768", "9890890890890")
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}
