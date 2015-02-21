package db

// В корне неправильный test-файл, зато пока он удовлетворяет потребности
// тестировать быстро меняющееся API.

import (
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"fmt"
	"testing"
	// "time"
)

func TestFunc(t *testing.T) {
	user1 := User{}
	user2 := User{
		Id:       2,
		Name:     "Aman",
		Password: "123",
	}

	err := user1.Load(1)
	fmt.Printf("err-load: %+v\n", err)
	fmt.Printf("user1: %+v\n\n", user1)

	err = user2.Save()
	fmt.Printf("err-save: %+v\n\n", err)

	all := []User{}
	Conn.Collection("User").Find(nil).All(&all)

	fmt.Printf("all: %+v\n\n", all)

	user2.Name = "Bman"
	err = user2.Update()
	fmt.Printf("err-upd: %+v\n\n", err)

	/*
		article := Article{
			Id:           1,
			User_id:      2,
			Title:        "conquer plan",
			Content:      "#TODO: develop a plan",
			CreationTime: time.Now(),
		}

		err = article.Save()
		fmt.Printf("err-a-save: %+v\n\n", err)
	*/

	article := Article{}
	err = article.Load(1)
	fmt.Printf("article: %+v\n", article)
	fmt.Printf("err-a-load: %+v\n\n", err)

	author, err := article.Author()
	fmt.Printf("author: %+v\n", author)
	fmt.Printf("err-a-author-load: %+v\n\n", err)
}
