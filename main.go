package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// This code demonstrates how to bind incoming JSON data to a struct.
// It uses the Gin framework's ShouldBindJSON function for this purpose.
// Additionally, it validates the data based on the rules defined in the struct tags.
type User struct {
	Name string `json:"name" binding:"required,max=100,min=5"`
	Age  int    `json:"age" binding:"required,max=18,min=5"`
}


func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"name": user.Name, "age": user.Age})
	})

	// Listen and serve on
    r.Run(":3000")
}