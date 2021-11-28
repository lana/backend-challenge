package handler

import (
	"encoding/json"
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

func TestBasketHandler(t *testing.T) {
	basketMock := models.Basket{
		Code:  "4200f350-4fa5-11ec-a386-1e003b1e5256",
		Items: make(map[string]models.Item),
		Total: 0,
	}

	gin.SetMode(gin.TestMode)

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("CreateBasket", mock.Anything, mock.Anything).Return(basketMock, nil)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.POST("/baskets", CreateBasketHandler(service))
		req, err := http.NewRequest(http.MethodPost, "/baskets", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("given a valid id request it returns 200", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(basketMock, nil)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.GET("/baskets/:id", GetBasketHandler(service))
		req, err := http.NewRequest(http.MethodGet, "/baskets/4200f350-4fa5-11ec-a386-1e003b1e5256", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("given a invalid params it returns 400", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(models.Basket{}, nil)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.GET("/baskets/:id", GetBasketHandler(service))
		req, err := http.NewRequest(http.MethodGet, "/baskets/", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a invalid id request it returns 400", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(models.Basket{}, models.ErrBasketNotFound)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.GET("/baskets/:id", GetBasketHandler(service))
		req, err := http.NewRequest(http.MethodGet, "/baskets/4200f350-4fa5-11ec-a386-1e003b1e5256", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

func TestRemoveBasketHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("given a empty basket id it returns 400", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("RemoveBasket", mock.Anything, mock.Anything).Return(models.ErrBasketNotFound)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.DELETE("/baskets/:id", RemoveBasketHandler(service))
		req, err := http.NewRequest(http.MethodDelete, "/baskets/", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a basket id it returns 200", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("RemoveBasket", mock.Anything, mock.Anything).Return(nil)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.DELETE("/baskets/:id", RemoveBasketHandler(service))
		req, err := http.NewRequest(http.MethodDelete, "/baskets/4200f350-4fa5-11ec-a386-1e003b1e5256", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("given a invalid basket id it returns 400", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("RemoveBasket", mock.Anything, mock.Anything).Return(models.ErrBasketNotFound)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.DELETE("/baskets/:id", RemoveBasketHandler(service))

		req, err := http.NewRequest(http.MethodDelete, "/baskets/4200f350-4fa5-11ec-a386-1e003b1e5256", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

func TestCheckoutBasketHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("given a empty basket id it returns 400", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.POST("/baskets/:id/checkout", CheckoutBasketHandler(service))

		url := fmt.Sprintf("/baskets/%s/checkout", "")
		req, err := http.NewRequest(http.MethodPost, url, nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a invalid basket id it returns 400", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(models.Basket{}, models.ErrBasketNotFound)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.POST("/baskets/:id/checkout", CheckoutBasketHandler(service))

		url := fmt.Sprintf("/baskets/%s/checkout", "667678678678")
		req, err := http.NewRequest(http.MethodPost, url, nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
	t.Run("given a basket id it returns 200", func(t *testing.T) {
		basketMock := models.Basket{
			Code: "4200f350-4fa5-11ec-a386-1e003b1e5256",
			Items: map[string]models.Item{
				"Tshirt": {
					Product: models.Product{
						Code:  "Tshirt",
						Name:  "Lana T-Shirt",
						Price: 20,
					},
					Quantity: 5,
					Total:    100,
				},
			},
			Total: 100,
		}

		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(basketMock, nil)

		basketMock2 := models.Basket{
			Code: "4200f350-4fa5-11ec-a386-1e003b1e5256",
			Items: map[string]models.Item{
				"Tshirt": {
					Product: models.Product{
						Code:  "Tshirt",
						Name:  "Lana T-Shirt",
						Price: 20,
					},
					Quantity: 5,
					Total:    75,
				},
			},
			Total: 75,
			Close: true,
		}
		repositoryMock.On("UpdateBasket", mock.Anything, mock.Anything).Return(basketMock2, nil)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.POST("/baskets/:id/checkout", CheckoutBasketHandler(service))

		url := fmt.Sprintf("/baskets/%s/checkout", "4200f350-4fa5-11ec-a386-1e003b1e5256")
		req, err := http.NewRequest(http.MethodPost, url, nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		var response models.Basket
		err = json.NewDecoder(res.Body).Decode(&response)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, basketMock2, response)
	})
}
