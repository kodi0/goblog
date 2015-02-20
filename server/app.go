package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.LoadHTMLGlob("templates/*")

	// api v1
	g.GET("/", func(c *gin.Context) {
		obj := gin.H{"title": "Home"}
		c.HTML(200, "index.tmpl", obj)
	})

	//static
	g.Static("/static", "./public")

	g.Run(":3000")
}
