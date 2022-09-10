package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/Vishal151/production-ready-api/internal/transport/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct{}

// Run - handles the startup of the application
func (app *App) Run() error {
	fmt.Println("Setting up our application")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to start server")
		return err
	}

	return nil
}

// Our main entrypoint for the application
func main() {
	fmt.Println("Go REST API course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up")
		fmt.Println(err)
	}
}
