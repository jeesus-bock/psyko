package main

import (
	"encoding/json"
	// "io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	log.Println("Setting up router.")

	r := gin.Default()

	r.GET("/tcl/:name", func(c *gin.Context) {
		// bodyAsByteArray, _ := io.ReadAll(c.Request.Body)
		bodyAsByteArray := []byte(`{ "nimi": "jussi", "ikä":42, "ammatti": "soturi" }`)
		jsonMap := make(map[string]interface{})
		json.Unmarshal(bodyAsByteArray, &jsonMap)

		scriptName := c.Params.ByName("name")
		output, err := HandleTcl("tcl/"+scriptName+".tcl", jsonMap)
		if err != nil {
			log.Fatal(err)
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
