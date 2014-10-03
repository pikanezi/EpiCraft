package controllers

import (
	"github.com/gocarina/pi"
	"EpiCraft/restful/models"
	"EpiCraft/restful/services"
)

// ConnectCharacterController connects the character to last place he was or the first zone if he never played.
func ConnectCharacterController(c *pi.RequestContext) error {
	account := c.Data["account"].(*models.PlayerAccount)
	idCharacter := c.GetRouteVariable("id")
	character, err := services.ConnectCharacterFromAccountByID(account, idCharacter)
	if err != nil {
		return pi.NewError(404, err)
	}
	return c.WriteDefault(character)
}

// CreateCharacterController creates a new character for the player.
func CreateCharacterController(c *pi.RequestContext) error {
	return nil
}
