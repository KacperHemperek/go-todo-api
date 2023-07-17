package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	controllers "github.com/kacperhemperek/go-todo-app/controllers"
)

func SetupRoutes(app *fiber.App, validate *validator.Validate) {

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos", controllers.CreateTodo)
	app.Get("/api/todos", controllers.GetTodos)
	app.Patch("/api/todos/:id/done", controllers.ToggleTodo)
	app.Delete("/api/todos/:id", controllers.DeleteTodo)
}
