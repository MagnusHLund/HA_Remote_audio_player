//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/yourusername/go-project/internal/app"
)

// InitializeApp creates and returns the application with all dependencies
func InitializeApp() (*app.App, error) {
	wire.Build(
		app.NewApp,
	)
	return nil, nil
}
