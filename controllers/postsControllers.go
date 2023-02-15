package controllers

import (
	"Go_Postgres_ORM/initializers"
	"Go_Postgres_ORM/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var postBody struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	bindError := c.Bind(&postBody)
	if bindError != nil {
		return
	}

	post := models.Post{Title: postBody.Title, Body: postBody.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func FindAllPosts(c *gin.Context) {

	var allPosts []models.Post

	initializers.DB.Find(&allPosts)

	c.JSON(200, gin.H{
		"posts": allPosts,
	})

}

func ShowSinglePost(c *gin.Context) {
	var singlePost models.Post

	id := c.Param("id")

	initializers.DB.First(&singlePost, id)

	c.JSON(200, gin.H{
		"post": singlePost,
	})
}

func UpdatePost(c *gin.Context) {
	var singlePostUpdate models.Post

	id := c.Param("id")

	var requestBody struct {
		Title string
		Body  string
	}
	c.Bind(&requestBody)

	initializers.DB.First(&singlePostUpdate, id)

	initializers.DB.Model(&singlePostUpdate).Updates(models.Post{Title: requestBody.Title, Body: requestBody.Body})

	c.JSON(200, gin.H{
		"Post Updated": singlePostUpdate,
	})

}

func DeletePosts(c *gin.Context) {
	var singlePostDelete models.Post

	id := c.Param("id")

	initializers.DB.First(&singlePostDelete, id)

	initializers.DB.Delete(&singlePostDelete)

	c.Status(200)

}
