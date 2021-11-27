package handler

import (
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

func TestHandler_Basket(t *testing.T) {
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

	t.Run("given a invalid id request it returns 200", func(t *testing.T) {
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
