package main

import (
	"Go_Postgres_ORM/controllers"
	"Go_Postgres_ORM/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/allPosts", controllers.FindAllPosts)
	r.GET("/myPost/:id", controllers.ShowSinglePost)
	r.PUT("/updatePost/:id", controllers.UpdatePost)
	r.DELETE("/deletePost/:id", controllers.DeletePosts)

	r.Run() // listen and serve on 0.0.0.0:8080
}
