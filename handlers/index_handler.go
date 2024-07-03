package handlers

import (
	"psyko/tcl_util"

	"github.com/gin-gonic/gin"
)

func IndexRequest(c *gin.Context) {
	endpoints := tcl_util.GetTclEndPoints()
	// Render the HTML file
	c.HTML(200, "alpine.html", gin.H{"endpoints": endpoints})
}
