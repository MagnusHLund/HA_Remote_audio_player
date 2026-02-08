package app

import "log"

// App represents the main application
type App struct{}

// NewApp creates a new App instance
func NewApp() *App {
	return &App{}
}

// Run executes the application
func (a *App) Run() error {
	log.Println("Application running")
	return nil
}
