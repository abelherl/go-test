package main

import (
	"github.com/abelherl/go-test/initializers"
	"github.com/abelherl/go-test/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
