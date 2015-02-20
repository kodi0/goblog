package main

import (
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
	g.Static("/static", "./server/public")
	//g.Use(static.Serve("/",static.LocalFile("")))

	g.Run(":3000")
}
