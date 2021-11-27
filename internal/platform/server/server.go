package server

import (
	"fmt"
	"log"
	"patriciabonaldy/lana/internal/lana"
	"patriciabonaldy/lana/internal/platform/server/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
	service  lana.Service
}

func New(port uint, service lana.Service) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf(":%d", port),
		service:  service,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", handler.CheckHandler())
	basket := s.engine.Group("/baskets")
	{
		basket.POST("", handler.CreateBasketHandler(s.service))
		basket.GET("/:id", handler.GetBasketHandler(s.service))
		basket.DELETE(":id", handler.RemoveBasketHandler(s.service))
		basket.POST("/:id/products", handler.AddProductHandler(s.service))
		basket.DELETE(":id/products", handler.RemoveProductHandler(s.service))
	}
}
