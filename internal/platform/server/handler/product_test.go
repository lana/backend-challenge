package handler

import (
	"bytes"
	"encoding/json"
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
	basketExpected := models.Basket{
		Code: "4200f350-4fa5-11ec-a386-1e003b1e5256",
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
	}

	gin.SetMode(gin.TestMode)

	t.Run("given a valid id request it returns 200", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("AddProduct", mock.Anything, mock.Anything, mock.Anything).Return(basketExpected, nil)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.POST("/baskets/products", AddProductHandler(service))

		body, _ := json.Marshal(request)
		reader := bytes.NewBuffer(body)
		req, err := http.NewRequest(http.MethodPost, "/baskets/products", reader)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("given a invalid request it returns 400", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("AddProduct", mock.Anything, mock.Anything, mock.Anything).Return(basketExpected, nil)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.POST("/baskets/products", AddProductHandler(service))

		body, _ := json.Marshal(lana.ProductRequest{})
		reader := bytes.NewBuffer(body)
		req, err := http.NewRequest(http.MethodPost, "/baskets/products", reader)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a invalid BasketID it returns 400", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("AddProduct", mock.Anything, mock.Anything, mock.Anything).
			Return(models.Basket{}, models.ErrBasketNotFound)

		service := lana.NewService(repositoryMock)
		r := gin.New()
		r.POST("/baskets/products", AddProductHandler(service))

		body, _ := json.Marshal(request)
		reader := bytes.NewBuffer(body)
		req, err := http.NewRequest(http.MethodPost, "/baskets/products", reader)
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
		repositoryMock.On("RemoveProduct", mock.Anything, mock.Anything, mock.Anything).Return(basketExpected, nil)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.DELETE("/baskets/products", RemoveProductHandler(service))

		body, _ := json.Marshal(request)
		reader := bytes.NewBuffer(body)
		req, err := http.NewRequest(http.MethodDelete, "/baskets/products", reader)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("given a invalid product inside request it returns 400", func(t *testing.T) {
		repositoryMock := new(storagemocks.Repository)
		repositoryMock.On("RemoveProduct", mock.Anything, mock.Anything, mock.Anything).Return(basketExpected, models.ErrProductNotFound)
		service := lana.NewService(repositoryMock)

		r := gin.New()
		r.DELETE("/baskets/products", RemoveProductHandler(service))

		body, _ := json.Marshal(request)
		reader := bytes.NewBuffer(body)
		req, err := http.NewRequest(http.MethodDelete, "/baskets/products", reader)
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
		r.DELETE("/baskets/products", RemoveProductHandler(service))

		body, _ := json.Marshal(lana.ProductRequest{})
		reader := bytes.NewBuffer(body)
		req, err := http.NewRequest(http.MethodDelete, "/baskets/products", reader)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}
