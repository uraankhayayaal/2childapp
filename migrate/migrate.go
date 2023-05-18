package main

import (
	"fmt"
	"log"

	"github.com/uraankhayayaal/2childapp/initializers"
	"github.com/uraankhayayaal/2childapp/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
	fmt.Println("? Migration complete")
}
