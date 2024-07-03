package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"psyko/db"
	"psyko/handlers"
)

func setup() *gin.Engine {
	err := db.SetupDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Setting up router.")
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	r.Static("/assets", "./assets")
	r.GET("/", handlers.IndexRequest)
	r.POST("/tcl/:name", handlers.TclRequest)
	return r
}

func main() {
	r := setup()
	r.Run("0.0.0.0:8080")
}
