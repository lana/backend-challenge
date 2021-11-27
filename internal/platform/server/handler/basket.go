package handler

import (
	"errors"
	"net/http"
	"patriciabonaldy/lana/internal/lana"
	"patriciabonaldy/lana/internal/models"

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
			switch {
			case errors.Is(err, models.ErrBasketNotFound):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.JSON(http.StatusOK, basket)
	}
}
