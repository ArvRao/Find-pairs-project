package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome!")
	})

	// Post request to find pairs of numbers summing to target
	app.Post("/find-pairs", func(c *fiber.Ctx) error {
		var requestBody RequestBody

		if err := c.BodyParser(&requestBody); err != nil {
			fmt.Println("An error occurred: ", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid input",
			})
		}

		// validation to check only integers present in numbers array
		if !validateArrayContainsIntegers(requestBody.Numbers) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Number array must have only integer numbers",
			})
		}

		// find pairs of numbers matching target value
		pairs := findAllPairs(requestBody.Numbers, int(requestBody.Target))

		return c.JSON(fiber.Map{
			"solutions": pairs,
		})
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))

}
