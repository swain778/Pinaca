package main

import (
	"fmt"
	controller "pinaca/Controller"
	database "pinaca/Database"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/books", controller.ReadBooks)
	r.POST("/book", controller.CreateBook)
	r.PUT("/books/:id", controller.UpdateBook)
	r.Run(":5000")
}
