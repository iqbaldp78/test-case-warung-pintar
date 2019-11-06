package main

import (
	"github.com/gin-gonic/gin"

	"test-warungpintar/controllers"
	"test-warungpintar/message"
	"test-warungpintar/tools"
)

func main() {
	// Set up a http setver.
	r := gin.Default()

	tools.TruncateFile()
	// Set blueprint request
	setBlueprintApp(r)
	setErrorHandler(r)
	r.Run(":8080")
}

/*setBlueprintApp used for handling request*/
func setBlueprintApp(router *gin.Engine) {
	router.GET("/sample", controllers.Sample)     // this endpoint used for generate response from server
	router.GET("/show", controllers.ShowResponse) // this endpoint used as client and request message to server
	router.GET("/ws", controllers.RunWs)          // this endpoint used for handling websocket
	router.GET("/index", controllers.Index)       // this endpoint used for serve html and will be trigger webscoket via browser
}

/*setErrorHandler used for handling 404 request*/
func setErrorHandler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		msg := message.New(0, nil)
		c.AbortWithStatusJSON(msg.StatusCode, msg)
	})
}
