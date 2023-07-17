package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	database "github.com/kacperhemperek/go-todo-app/db"
	"github.com/kacperhemperek/go-todo-app/initializers"
	"github.com/kacperhemperek/go-todo-app/router"
)

func init() {
	initializers.LoadEnv()
	database.InitializeDb()
}

func main() {

	validate := validator.New()

	port := fmt.Sprintf(":%s", initializers.ServerSecrets.PORT)

	app := fiber.New()

	app.Use(cors.New())

	router.SetupRoutes(app, validate)

	log.Fatal(app.Listen(port))
}
