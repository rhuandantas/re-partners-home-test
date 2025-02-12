// wire.go
//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rhuandantas/re-partners-home-test/internal/adapters/http"
	"github.com/rhuandantas/re-partners-home-test/internal/adapters/http/handlers"
	"github.com/rhuandantas/re-partners-home-test/internal/adapters/repository/cache"
	"github.com/rhuandantas/re-partners-home-test/internal/core/usecases"
)

func InitializeWebServer() (*http.Server, error) {
	wire.Build(cache.NewMemcacheClient,
		usecases.NewStorePackSize,
		usecases.NewPackItem,
		handlers.NewPackHandler,
		http.NewAPIServer)
	return &http.Server{}, nil
}
