package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	validator, err := NewValidator()
	if err != nil {
		log.Fatal(err)
	}

	router.POST("/validator", validator.Handler)
	router.Run(":8080")
}
