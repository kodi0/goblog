package main

import (
	//"fmt"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type Article struct {
	Id      bson.ObjectId `bson:"_id"`
	Title   string        `json:"title"`
	Subject string        `json:"subject"`
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
	// Api root
	api := router.Group("/api/v1")

	// GET All articles /api/v1/articles
	api.GET("/articles", func(c *gin.Context) {

		articles_result := []Article{}

		err = articles.Find(bson.M{}).All(&articles_result)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(articles_result)
	})
	// Show article
	api.GET("/article/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")

		oid := bson.ObjectIdHex(id) // convert to ObjectId
		article_result := Article{}

		err = articles.FindId(oid).One(&article_result)
		if err != nil {
			log.Printf("no document found %v\n", err)
		}
		fmt.Println(article_result)
	})

	// Create article POST /api/v1/articles
	api.POST("/articles", func(c *gin.Context) {
		var json Article
		c.Bind(&json)

		// Insert document into collection Articles
		err = articles.Insert(&Article{bson.NewObjectId(), json.Title, json.Subject})
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
