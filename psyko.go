package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	log.Println("Setting up router.")

	r := gin.Default()

	r.POST("/tcl/:name", func(c *gin.Context) {
		bodyAsByteArray, _ := io.ReadAll(c.Request.Body)
		jsonMap := make(map[string]interface{})
		json.Unmarshal(bodyAsByteArray, &jsonMap)

		scriptName := c.Params.ByName("name")

		script, err := os.ReadFile("scripts/" + scriptName + ".tcl")
		if err != nil {
			log.Fatal(err.Error())
		}
		output, err := HandleTcl(string(script), jsonMap)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("output: ", output)
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte(output))
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run("0.0.0.0:8080")
}
