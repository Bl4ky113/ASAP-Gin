
package app

import (
	"github.com/gin-gonic/gin"
	"asap/app/teacher"
)

var app *gin.Engine = nil

func StartApp () *gin.Engine {
	app = gin.Default()

	addGroupsToRouter()

	return app;
}

func addGroupsToRouter () int {
	var teacherGroup *gin.RouterGroup = app.Group("/teacher")
	teacher.AddTeacherRoutes(teacherGroup)

	return 0;
}
