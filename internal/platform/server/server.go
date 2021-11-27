package server

import (
	"fmt"
	"log"
	"patriciabonaldy/lana/internal/lana"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
	// deps
	contactService lana.Service
}

func New(port uint, service lana.Service) Server {
	srv := Server{
		engine:         gin.New(),
		httpAddr:       fmt.Sprintf(":%d", port),
		contactService: service,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	// TODO: create routes
}
