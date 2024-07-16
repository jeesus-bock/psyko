package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"psyko/db"
	"psyko/handlers"
	"psyko/tcl_util"
)

func setup() *gin.Engine {
	// Run the database initialization routine.
	err := db.SetupDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Run the TCL initialization routine.
	err = tcl_util.SetupTcl()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Setting up router.")
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
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
