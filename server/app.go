package main

import (
	// log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	// "github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "net/http"
	"time"
)

// Post - структура описывающая публикацию для блога
type Post struct {
	ID        bson.ObjectId `bson:"_id, omitempty"`
	Title     string        `json:"title" binding:"required"`
	Body      string        `json:"body" binding:"required"`
	AuthorID  bson.ObjectId
	CreatedAt time.Time
	UpdatedAt time.Time
}

func sessionDB() *mgo.Session {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	defer session.Close()

	return session
}

var sess *mgo.Session

//
func main() {
	router := gin.Default()

	// Все роутеры
	sess = sessionDB()

	router.POST("/article", articleAdd)
	router.GET("/article/:id", articleGet)
	router.PUT("/article", articleUpdate)
	router.DELETE("/article/:id", articleDelete)
	router.GET("/articles", articlesGetAll)

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

// Добавление статьи
func articleAdd(c *gin.Context) {
	var post Post

	c.Bind(&post)
	post.AuthorID = bson.NewObjectId() // временное решение, потом будет подставляться правильное Id текущего автора
	post.CreatedAt = time.Now()

	// сохранить в БД
	article := sess.DB("test").C("article")

	// Insert Datas
	err := article.Insert(post)

	if err != nil {
		panic(err)
	}

}

// Извлечение статьи по id
func articleGet(c *gin.Context) {
	c.String(200, "pong")
}

// Редактирование статьи
func articleUpdate(c *gin.Context) {
	c.String(200, "pong")
}

// Удаление статьи
func articleDelete(c *gin.Context) {
	c.String(200, "pong")
}

// Вывод всех статей с кратким содержанием статьи (140 символов)
func articlesGetAll(c *gin.Context) {
	c.String(200, "pong")
}

// Выводить статьи на главную

// У статьи должен быть автор и время публикации

// Рабочий кабинет со статьями автора

// Регистрация

// Авторизация
