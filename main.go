package main

import (
	"fmt"
	"log"
	"time"
    "math/rand/v2"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		Concurrency: 100,
	})

	app.Get("/", func(c fiber.Ctx) error {
		fmt.Println("start sleeping...")
		sleep_seconds := rand.IntN(8) + 1
		time.Sleep(time.Duration(sleep_seconds) * time.Second)
		fmt.Printf("done sleeping after %d seconds!\n", sleep_seconds)
		// Send a string response to the client
        msg := fmt.Sprintf("Hello, World 👋!, Slept %d seconds", sleep_seconds)
		return c.SendString(msg)
	})

	log.Fatal(app.Listen(":8000"))
}
