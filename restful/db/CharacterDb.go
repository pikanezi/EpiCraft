package db

import (
	"EpiCraft/restful/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	characterCollectionName = "characters"
)
// GetCharacterByIDAndAccount finds the character by its account login and character id.
func GetCharacterByIDAndAccount(login, idCharacter string) (*models.Character, error) {
	return nil, nil
}

// CreateCharacter creates a new character for the account.
func CreateCharacter(account *models.PlayerAccount, character *models.Character) (*models.Character, error) {
	character.LoginAccount = account.Login
	character.ID = bson.NewObjectId()
	change := mgo.Change{
		Update: character,
		Upsert: true,
		ReturnNew: true,
	}
	_, err := c(characterCollectionName).Find(character).Apply(change, &character)
	return character, err
}
