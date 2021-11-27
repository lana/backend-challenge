package handler

import (
	"net/http"
	"patriciabonaldy/lana/internal/lana"

	"github.com/gin-gonic/gin"
)

func AddProductHandler(service lana.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req lana.ProductRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		basket, err := service.AddProduct(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, basket)
	}
}

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
