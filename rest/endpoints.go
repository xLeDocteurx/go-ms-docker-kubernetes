package main

import (
	// "fmt"
	"github.com/gofiber/fiber/v2"
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "io/ioutil"
	// "encoding/json"
	// "log"
	// "net/http"
)

func GetPeople(c *fiber.Ctx) error {
    return c.JSON(People)
}
func GetPersonByID(c *fiber.Ctx) error {
    id := c.Params("id")
    for _, a := range People {
        if a.ID == id {
            return c.JSON(a)
            // return
        }
    }
    return c.SendString("person not found")
}
func PostPerson(c *fiber.Ctx) error {
    var newPerson Person
    People = append(People, newPerson)
    return c.JSON(newPerson)
}
func RemovePersonById(c *fiber.Ctx) error {
    id := c.Params("id")
    People = FilterPeople(People, func(person Person) bool {
		return person.ID == id
	})
    return c.SendString(id)
}

func GetPlanets(c *fiber.Ctx) error {
    return c.JSON(Planets)
}
func GetPlanetByID(c *fiber.Ctx) error {
    id := c.Params("id")
    for _, a := range Planets {
        if a.ID == id {
            return c.JSON(a)
        }
    }
    return c.SendString("planet not found")
}
func PostPlanet(c *fiber.Ctx) error {
    var newPlanet Planet
    Planets = append(Planets, newPlanet)
    return c.JSON(newPlanet)
}
func RemovePlanetById(c *fiber.Ctx) error {
    id := c.Params("id")
    Planets = FilterPlanets(Planets, func(planet Planet) bool {
		return planet.ID == id
	})
    return c.SendString(id)
}
