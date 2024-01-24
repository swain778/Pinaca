package controller

import (
	"fmt"
	database "pinaca/Database"
	"pinaca/model"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book model.Book

	res := database.DB.Create(book)
	if res.RowsAffected == 0 {
		fmt.Println("Error while creating a book data")
	}
}

func ReadBooks(c *gin.Context) {
	var book []model.Book

	res := database.DB.Find(book)
	if res.Error != nil {
		fmt.Println("Error while fetching all book details")
	}
}

func UpdateBook(c *gin.Context) {
	var book model.Book
	id := c.Param("id")
	res := database.DB.Model(&book).Where("id = ?", id).Updates(book)
	if res.RowsAffected == 0 {
		fmt.Println("Error while updating Book record")
	}
}
