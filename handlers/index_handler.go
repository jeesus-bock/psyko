package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"psyko/tcl_util"

	"github.com/gin-gonic/gin"
)

func IndexRequest(c *gin.Context) {
	endpoints := tcl_util.GetTclEndPoints()
	endpointJson, err := json.Marshal(endpoints)
	log.Printf("endpointJson: %+v", endpointJson)
	if err != nil {
		c.Writer.WriteHeader(http.StatusExpectationFailed)
		c.Writer.Write([]byte(err.Error()))
	}
	// Render the HTML file
	c.HTML(200, "index.html", gin.H{"endpoints": endpoints})
}
