package routes

import (
	"net/http"

	"github.com/Amartya-Bhardwaj/goAPI/models"
	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

//@req GET request
//@desc Get all books
func GetBook(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(200, gin.H{"data": books})
}

//POST request
//@desc Create a new entry
func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

//PUT request
//@desc Update a existing record
func UpdateBook(c *gin.Context){
	//get model
	var book models.Book
	if err := models.DB.Where("id = ?",c.Param("id")).First(&book).Error; err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Book not found!",
		})
		return
	}
	//input validation
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	models.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

//DELETE request
//@desc Delete a record

func DeleteBook(c *gin.Context){
	//get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?",c.Param("id")).First(&book).Error; err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Book not found!",
		})
		return
	}
	models.DB.Delete(&book)
	c.JSON(200,gin.H{
		"msg": "Deleted",
	})
}