package main

import (
	"fmt"
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
		return c.SendString("Hello, World !")
	})

	// post todo
	app.Post("/api/todos", func(c fiber.Ctx) error {
		todo := &Todo{}

		// Парсим тело запроса в структуру Todo с помощью Bind
		if err := c.Bind().Body(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Обработка ответа если пустая строка
		if todo.Title == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Title is required",
			})
		}

		// Проверяем, существует ли задача с таким же заголовком
		for _, t := range todos {
			if t.Title == todo.Title {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"error": "Задача с таким заголовком уже существует",
				})
			}
		}

		// Добавляем новую задачу в список
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		// Возвращаем созданную задачу в ответе
		return c.Status(201).JSON(todo)
	})

	// update checkBox
	app.Patch("/api/todos/:id", func(c fiber.Ctx) error {
		// Вытаскиваем id из параметров
		id := c.Params("id")

		// Поиск в слайсе todos нужный элемент
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	})

	// get all
	app.Get("/api/todos/getAll", func(c fiber.Ctx) error {
		return c.JSON(todos)
	})
	log.Fatal(app.Listen(":3000"))
}
