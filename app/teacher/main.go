
package teacher

import (
	"github.com/gin-gonic/gin"
)

func AddTeacherRoutes (group *gin.RouterGroup) int {
	group.GET("/", getTeachers)
	return 0;
}
