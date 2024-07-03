package handlers

import (
	"github.com/gin-gonic/gin"
	"psyko/tcl_util"
)

func IndexRequest(c *gin.Context) {
	endpoints := tcl_util.GetTclEndPoints()
	// Render the HTML file
	c.HTML(200, "index.html", endpoints)
}
