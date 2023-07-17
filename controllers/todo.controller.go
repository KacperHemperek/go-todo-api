package controllers

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	database "github.com/kacperhemperek/go-todo-app/db"
	"github.com/kacperhemperek/go-todo-app/models"
)

func CreateTodo(c *fiber.Ctx) error {
	validate := validator.New()

	todo := models.Todo{}

	todo.Done = false

	if err := c.BodyParser(&todo); err != nil {
		errorMessage := fmt.Sprintf("Couldn't create a todo, cause: %s", err.Error())

		return c.Status(400).SendString(errorMessage)
	}

	err := validate.Struct(todo)

	if err != nil {
		var errors []models.ValidationError

		for _, err := range err.(validator.ValidationErrors) {
			validationError := models.ValidationError{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			}

			errors = append(errors, validationError)
		}

		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid user input",
			"errors":  errors,
		})
	}

	database.DB.Create(&todo)

	return c.JSON(todo)
}

func GetTodos(c *fiber.Ctx) error {

	var todos []models.Todo

	database.DB.Order("created_at desc").Find(&todos)

	return c.JSON(todos)
}

func ToggleTodo(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 0)

	if err != nil {
		return c.Status(401).SendString("Invalid ID - ID must be a number")
	}

	var todo models.Todo

	queryErr := database.DB.First(&todo, id).Error

	if queryErr != nil {
		errorMessage := fmt.Sprintf("Couldn't find a todo with id: %d", id)

		return c.Status(404).JSON(fiber.Map{
			"message": errorMessage,
		})
	}

	todo.Done = !todo.Done
	saveErr := database.DB.Save(&todo).Error

	if saveErr != nil {
		errorMessage := fmt.Sprintf("Couldn't update a todo with id: %d", id)

		return c.Status(500).JSON(fiber.Map{
			"message": errorMessage,
		})
	}

	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 0)

	if err != nil {
		return c.Status(401).SendString("Invalid ID - ID must be a number")
	}

	var todo models.Todo

	queryErr := database.DB.First(&todo, id).Error

	if queryErr != nil {
		errorMessage := fmt.Sprintf("Couldn't find a todo with id: %d", id)

		return c.Status(404).JSON(fiber.Map{
			"message": errorMessage,
		})
	}

	deleteErr := database.DB.Delete(&todo).Error

	if deleteErr != nil {
		errorMessage := fmt.Sprintf("Couldn't delete a todo with id: %d", id)

		return c.Status(500).JSON(fiber.Map{
			"message": errorMessage,
		})
	}

	return c.Status(200).JSON(todo)

}
