package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	// Define a route for the GET method on the root path '/'

	// Start the server on port 3000
	log.Println("Server is running on port 8080")

	err := app.Listen(":8080")
	if err != nil {
		return
	}

}
