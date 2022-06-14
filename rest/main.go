package main

import (
	// "fmt"
	// "github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	// "net/http"
	// "io/ioutil"
	// "encoding/json"
	"log"
	// "net/http"
)

func setUp() {
    app := fiber.New()
    app.Get("/people", GetPeople)
    app.Get("/people/:id", GetPersonByID)
    app.Post("/people", PostPerson)
	
    app.Get("/planets", GetPlanets)
    app.Get("/planets/:id", GetPlanetByID)
    app.Post("/planets", PostPlanet)

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the awesome Star Wars API")
	})

    app.Listen(":3000")
}

func loop() {
}

func main() {
	log.Println("Starting...")

	setUp()

	go func() {
		for {
			loop()
		}
	}()

	// // Executed after app exit
	// defer fmt.Println("This app is Exiting...")
}
