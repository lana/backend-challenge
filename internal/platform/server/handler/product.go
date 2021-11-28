package handler

import (
	"net/http"
	"patriciabonaldy/lana/internal/lana"

	"github.com/gin-gonic/gin"
)

// AddProductHandler add a new product to basket.
// return 201 if this could be created.
// Otherwise, it will return 500
func AddProductHandler(service lana.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req lana.ProductRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		_, err := service.AddProduct(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}

// RemoveProductHandler remove a product inside a basket.
// require a basket id and product id.
// it will return 200 if this is ok.
// otherwise will return 400
func RemoveProductHandler(service lana.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req lana.ProductRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		basket, err := service.RemoveProduct(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, basket)
	}
}
