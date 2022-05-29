package main

import (
	"github.com/Amartya-Bhardwaj/goAPI/models"
	"github.com/Amartya-Bhardwaj/goAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books",routes.GetBook)
	r.POST("/",routes.CreateBook)
	r.PUT("/:id",routes.UpdateBook)
	r.DELETE("/:id",routes.DeleteBook)
	r.Run()
}
