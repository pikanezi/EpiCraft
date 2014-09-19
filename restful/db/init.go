package db

import (
	"github.com/pikanezi/tokauth"
	mgo "gopkg.in/mgo.v2"
	"time"
)

const (
	databaseName         = "epicraft"
	accessCollectionName = "access"
)

var (
	mongoSession *mgo.Session
	mongoPool    = make(map[int]*mgo.Session)
)

func init() {
	var err error
	mongoSession, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	mongoSession.SetSafe(&mgo.Safe{})
	accountCollection := mongoSession.DB(databaseName).C(accountCollectionName)
	accountCollection.DropIndex("login")
	index := mgo.Index{
		Key:    []string{"login"},
		Unique: true,
	}
	if err := accountCollection.EnsureIndex(index); err != nil {
		panic(err)
	}
	tokauth.SetTokenExpiration(30 * time.Minute)
	tokauth.SetAccessCollection(mongoSession.DB(databaseName).C(accessCollectionName))
	tokauth.SetClientCollection(accountCollection)
}

// getSession returns the first available session.
func getSession() *mgo.Session {
	for _, session := range mongoPool {
		return session
	}
	return nil
}

// c returns the collection asked.
func c(coll string) *mgo.Collection {
	return getSession().DB(databaseName).C(coll)
}

// Copy a session of MongoDB.
func Copy(seed int) {
	mongoPool[seed] = mongoSession.Copy()
}

// CloseCopy closes a copy of a MongoDB session.
func CloseCopy(seed int) {
	mongoPool[seed].Close()
	delete(mongoPool, seed)
}
