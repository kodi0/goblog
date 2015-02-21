package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Indexer interface {
	Index() mgo.Index
}

type ORM interface {
	BsonId() bson.M

	Load() error
	Save() error
	Update() error
	Delete() error
}

func uniqIndex(key []string) mgo.Index {
	return mgo.Index{
		Key:        key,
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

/*
 * User:
 */

type User struct {
	Id uint64

	Name     string
	Password string
}

func (_ User) Index() mgo.Index {
	return uniqIndex([]string{"id"})
}

func (u *User) BsonId() bson.M {
	return bson.M{"id": u.Id}
}

func (u *User) Load(id uint64) error {
	return Conn.Collection("User").Find(bson.M{"id": id}).Limit(1).One(u)
}

func (u *User) Save() error {
	return Conn.Collection("User").Insert(u)
}

func (u *User) Update() error {
	return Conn.Collection("User").Update(u.BsonId(), &u)
}

func (u *User) Delete() error {
	return Conn.Collection("User").Remove(u.BsonId())
}

/*
 * Article:
 */

type Article struct {
	Id      uint64
	User_id uint64

	Title   string
	Content string

	CreationTime time.Time
}

func (_ Article) Index() mgo.Index {
	return uniqIndex([]string{"id", "title"})
}

func (a *Article) BsonId() bson.M {
	return bson.M{"id": a.Id}
}

func (a *Article) Load(id uint64) error {
	return Conn.Collection("Article").Find(bson.M{"id": id}).Limit(1).One(a)
}

func (a *Article) Save() error {
	return Conn.Collection("Article").Insert(a)
}

func (a *Article) Update() error {
	return Conn.Collection("Article").Update(a.BsonId(), &a)
}

func (a *Article) Delete() error {
	return Conn.Collection("Article").Remove(a.BsonId())
}

func (a *Article) Author() (*User, error) {
	author := &User{}

	if err := author.Load(a.User_id); err != nil {
		return nil, err
	}

	return author, nil
}
