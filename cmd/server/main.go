package main

import (
	"fmt"
)

// App - the struct which contains things like pointers
// to database connections
type App struct{}

// Run - handles the startup of the application
func (app *App) Run() error {
	fmt.Println("Setting up our application")
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
