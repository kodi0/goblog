package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
	"net/http"
)

//
func main() {
	//IsDevelopment recompiles template on reload
	r := render.New(render.Options{
		IsDevelopment: true,
	})
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		r.HTML(c.Writer, http.StatusOK, "index", nil)
	})

	router.Static("static", "./public")

	// Native implementation
	//fs := http.FileServer(http.Dir("public"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	router.Run(":3000")
}
