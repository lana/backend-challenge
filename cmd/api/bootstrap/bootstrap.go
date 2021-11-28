package bootstrap

import (
	"patriciabonaldy/lana/internal/lana"
	"patriciabonaldy/lana/internal/platform/server"
	"patriciabonaldy/lana/internal/platform/storage/platform/storage/memory"
)

const (
	port = 8080
)

// Run application
func Run() error {
	repository := memory.NewRepository()
	lanaService := lana.NewService(repository)
	srv := server.New(port, lanaService)
	return srv.Run()
}
