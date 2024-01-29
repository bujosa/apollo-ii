package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Demos the binding of JSON data to a struct for validation
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

    r.Run() // listen and serve on 0.0.0.0:8080
}