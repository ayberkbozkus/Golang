package main

import (
    "github.com/gofiber/fiber/v2"
    "fmt"
)

func helloWorld(c *fiber.Ctx) error {
    return c.SendString("Hello, World 👋!")
}

func apiFunc(c *fiber.Ctx) error {
    msg := fmt.Sprintf("✋ %s", c.Params("*"))
    return c.SendString(msg) // => ✋ register
}

func main() {
    app := fiber.New()

    app.Get("/", helloWorld)

    app.Get("/api/*", apiFunc)

    app.Listen(":3000")
}