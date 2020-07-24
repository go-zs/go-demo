package db

import "gopkg.in/mgo.v2"

type (
	ModelManager struct {
		Session  *mgo.Session
		DataBase string
	}
)
