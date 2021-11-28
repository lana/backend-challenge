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
		id := ctx.Param("id")
		if id == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}

		code := ctx.Param("code")
		if code == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}
		req := lana.ProductRequest{
			BasketID:    id,
			ProductCode: code,
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
		id := ctx.Param("id")
		if id == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}

		code := ctx.Param("code")
		if code == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}
		req := lana.ProductRequest{
			BasketID:    id,
			ProductCode: code,
		}
		basket, err := service.RemoveProduct(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, basket)
	}
}
