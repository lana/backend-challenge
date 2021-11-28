package handler

import (
	"net/http"
	"patriciabonaldy/lana/internal/lana"

	"github.com/gin-gonic/gin"
)

// CreateBasketHandler create a new basket.
// return 201 if this could be created.
// Otherwise, it will return 500
func CreateBasketHandler(service lana.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		basket, err := service.CreateBasket(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, basket)
	}
}

// GetBasketHandler return basket.
// require a basket id and
// return 200 if this is ok.
// otherwise will return 400
func GetBasketHandler(service lana.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}

		basket, err := service.GetBasket(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}

		ctx.JSON(http.StatusOK, basket)
	}
}

// RemoveBasketHandler remove a basket.
// require a basket id.
// it will return 200 if this is ok.
// otherwise will return 400
func RemoveBasketHandler(service lana.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}

		err := service.RemoveBasket(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}

		ctx.Status(http.StatusOK)
	}
}

// CheckoutBasketHandler close a basket.
// require a basket id.
// it will return 200 if this is ok.
// otherwise will return 400
func CheckoutBasketHandler(service lana.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}

		basket, err := service.CheckoutBasket(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, basket)
	}
}
