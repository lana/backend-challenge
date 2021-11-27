package bootstrap

import (
	"patriciabonaldy/lana/internal/lana"
	"patriciabonaldy/lana/internal/platform/server"
)

const (
	port = 8080
)

// Run application
func Run() error {
	lanaService := lana.NewService()
	srv := server.New(port, lanaService)
	return srv.Run()
}
