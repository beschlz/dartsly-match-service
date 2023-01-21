package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterStatusEndpoints(r *gin.Engine) {
	r.GET("/", getStatus)
}

func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "running",
	})
}
