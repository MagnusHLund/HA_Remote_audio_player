package app

import "github.com/google/wire"

// App represents the main application.
type App struct {
	controller *Controller
}

// ProviderSet wires the app-level dependencies.
var ProviderSet = wire.NewSet(
	NewLogger,
	NewController,
	NewApp,
)

// NewApp creates a new App instance.
func NewApp(controller *Controller) *App {
	return &App{controller: controller}
}

// Run executes the application.
func (a *App) Run() error {
	return a.controller.Run()
}
