package models

import (
	"EpiCraft/game/entities"
	"gopkg.in/mgo.v2/bson"
)

// Character represents a single character of a playerAccount.
type Character struct {
	*entities.Character `bson:",inline"`

	ID           bson.ObjectId `bson:"_id" json:"id,omitempty"`
	LoginAccount string        `bson:"loginAccount" json:"-"`
}
