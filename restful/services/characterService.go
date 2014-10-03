package services

import (
	"EpiCraft/restful/models"
	"EpiCraft/restful/db"
	"EpiCraft/game"
)

// ConnectCharacterFromUserByID connects the character of the user.
func ConnectCharacterFromAccountByID(user *models.PlayerAccount, idCharacter string) (*models.Character, error){
	character, err := db.GetCharacterByIDAndAccount(user.Login, idCharacter)
	if err != nil {
		return nil, err
	}
	game.GetWorld().AddPlayer(character.Character)
	return character, nil
}

// CreateCharacter creates a character for the account.
func CreateCharacter(account *models.PlayerAccount, character *models.Character)  (*models.Character, error) {
	return db.CreateCharacter(account, character)
}
