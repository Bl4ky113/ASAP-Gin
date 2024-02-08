
package teacher

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func getTeachers (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
	return
}
