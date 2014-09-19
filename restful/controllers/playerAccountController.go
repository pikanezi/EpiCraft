package controllers

import (
	"EpiCraft/restful/db"
	"EpiCraft/restful/models"
	"github.com/gocarina/pi"
)

// RegisterAccountController registers the user.
func RegisterAccountController(c *pi.RequestContext) error {
	account := &models.PlayerAccount{}
	if err := c.GetDefaultObject(account); err != nil {
		return pi.NewError(400, err)
	}
	account, err := db.CreatePlayerAccount(account)
	if err != nil {
		return pi.NewError(400, err)
	}
	return c.WriteDefault(account.WithRefreshToken())
}

// RefreshAccountController refreshes the token of the user.
func RefreshAccountController(c *pi.RequestContext) error {
	account := &models.PlayerAccount{}
	if err := c.GetDefaultObject(account); err != nil {
		return pi.NewError(400, err)
	}
	account, err := db.RefresPlayerAccount(account)
	if err != nil {
		return pi.NewError(400, err)
	}
	return c.WriteDefault(account.Secure())
}
