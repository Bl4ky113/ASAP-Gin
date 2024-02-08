
package main

import (
	"github.com/gin-gonic/gin"
	"asap/app"
)


func main () {
	var AsapApp *gin.Engine = app.StartApp()

	AsapApp.Run(":8080")
}
