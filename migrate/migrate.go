package main

import (
	"github.com/roadleft/roadleft-msgc/initializers"
	"github.com/roadleft/roadleft-msgc/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()

}

func main() {

	initializers.DB.AutoMigrate(&models.User{})
}
