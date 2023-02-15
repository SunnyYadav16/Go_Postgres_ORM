package main

import (
	"Go_Postgres_ORM/initializers"
	"Go_Postgres_ORM/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
