package db

import (
	"cfg"
	"gopkg.in/mgo.v2"
)

type Connection struct {
	host, dbName string

	sess *mgo.Session
}

var (
	Conn *Connection
)

func init() {
	var sess *mgo.Session

	sess, err := mgo.Dial(cfg.HOST)
	if err != nil {
		panic(err)
	}

	Conn = &Connection{
		dbName: cfg.DB,
		sess:   sess,
	}
}

func (conn *Connection) Collection(name string) *mgo.Collection {
	return conn.sess.DB(conn.host).C(name)
}

func (conn *Connection) Session() *mgo.Session {
	return conn.sess.Copy()
}

func (conn *Connection) Close() {
	conn.sess.Close()
}

func (conn *Connection) IndexedCollection(row Indexer, name string) (*mgo.Collection, error) {
	collection := conn.Collection(name)

	if err := collection.EnsureIndex(row.Index()); err != nil {
		return nil, err
	}

	return collection, nil
}
