package database

import (
	"log"

	"github.com/kacperhemperek/go-todo-app/initializers"
	"github.com/kacperhemperek/go-todo-app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDb() {
	var err error

	DB, err = gorm.Open(sqlite.Open(initializers.ServerSecrets.DATABASE_URL))

	if err != nil {
		log.Fatal(err.Error())
	}

	DB.AutoMigrate(&models.Todo{})
	DB.AutoMigrate(&models.User{})

	log.Println("Migration applied and database connection initialized")
}
