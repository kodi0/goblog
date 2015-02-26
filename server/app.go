package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Subject string `json:"subject"`
}

func main() {

	// Server stuff
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Define articles collection
	articles := session.DB("goblog").C("Articles")

	//IsDevelopment recompiles template on reload
	r := render.New(render.Options{
		IsDevelopment: true,
	})
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		r.HTML(c.Writer, http.StatusOK, "index", nil)
	})

	api := router.Group("/v1")

	api.POST("/articles", func(c *gin.Context) {
		var json Article
		c.Bind(&json)

		// Insert document into collection Articles
		err = articles.Insert(&Article{json.Title, json.Subject})
		if err != nil {
			log.Fatal(err)
		}

	})

	router.Static("static", "./public")

	// Native implementation
	//fs := http.FileServer(http.Dir("public"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	router.Run(":3000")
}
