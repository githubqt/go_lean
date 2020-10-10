package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		data := map[string]interface{}{
			"lag": "go",
			"tag": "<br>",
			"tagd": "<php>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})
	r.Run(":8082") // listen and serve on 0.0.0.0:8080
}
