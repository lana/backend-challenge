package handler

import (
	"net/http"
	"patriciabonaldy/lana/internal/lana"
	"strings"

	"github.com/gin-gonic/gin"
)

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

func RemoveBasketHandler(service lana.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if strings.Trim(id, " ") == "" {
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
