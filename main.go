package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/api/todos", func(c fiber.Ctx) error {
		var todo Todo

		// Парсим тело запроса в структуру Todo с помощью Bind
		if err := c.Bind().Body(&todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Обработка ответа если пустая строка
		if todo.Title == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Title is required",
			})
		}

		// Добавляем новую задачу в список
		todo.ID = len(todos) + 1
		todos = append(todos, todo)

		// Возвращаем созданную задачу в ответе
		return c.Status(201).JSON(todo)
	})

	log.Fatal(app.Listen(":3000"))
}
