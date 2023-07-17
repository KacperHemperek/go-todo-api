package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Secrets struct {
	JWT_SECRET   string `validate:"required"`
	JWT_PUBLIC   string `validate:"required"`
	DATABASE_URL string `validate:"required"`
	PORT         string `validate:"required"`
}

var ServerSecrets Secrets

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
	}

	ServerSecrets = Secrets{
		JWT_SECRET:   os.Getenv("JWT_SECRET"),
		DATABASE_URL: os.Getenv("DATABASE_URL"),
		JWT_PUBLIC:   os.Getenv("JWT_PUBLIC"),
		PORT:         os.Getenv("PORT"),
	}

	validate := validator.New()

	err = validate.Struct(ServerSecrets)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("Loaded env variables")
}
