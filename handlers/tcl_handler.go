package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"psyko/tcl_util"
)

func TclRequest(c *gin.Context) {
	bodyAsByteArray, _ := io.ReadAll(c.Request.Body)
	jsonMap := make(map[string]interface{})
	json.Unmarshal(bodyAsByteArray, &jsonMap)

	scriptName := c.Params.ByName("name")
	if scriptName == "helpers" {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		return
	}

	// Load up the helper procedures for scripts to use.
	helpers, err := os.ReadFile("scripts/helpers.tcl")
	if err != nil {
		log.Fatal(err.Error())
	}

	script, err := os.ReadFile("scripts/" + scriptName + ".tcl")
	if err != nil {
		log.Fatal(err.Error())
	}
	output, err := tcl_util.RunTcl(string(helpers)+string(script), jsonMap)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("output: ", output)
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(output))
}
