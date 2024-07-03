package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"psyko/handlers"
)

func setupRouter() *gin.Engine {
	log.Println("Setting up router.")

	r := gin.Default()

	r.POST("/tcl/:name", handlers.HandleTclRequest)
	return r
}

func main() {
	r := setupRouter()
	r.Run("0.0.0.0:8080")
}
