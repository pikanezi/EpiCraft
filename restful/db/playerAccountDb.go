package db

import (
	"EpiCraft/restful/models"
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"github.com/pikanezi/tokauth"
	"time"
	"gopkg.in/mgo.v2/bson"
)

const (
	accountCollectionName = "accounts"
)

// CreatePlayerAccount creates the account if it does not already exists.
func CreatePlayerAccount(account *models.PlayerAccount) (*models.PlayerAccount, error) {
	if err := account.ValidateAndCryptPassword(); err != nil {
		return nil, err
	}
	err := c(accountCollectionName).Insert(account)
	if err != nil && err.(*mgo.LastError).Code == 11000 {
		return nil, fmt.Errorf(`login [%s] already taken`, account.Login)
	}
	refreshToken, accessToken, err := tokauth.Register(account.Login)
	if err != nil {
		return nil, err
	}
	account.AccessToken = accessToken
	account.RefreshToken = refreshToken
	return account, nil
}

// RefresPlayerAccount refreshes the token of the account.
func RefresPlayerAccount(account *models.PlayerAccount) (*models.PlayerAccount, error) {
	accessToken, err := tokauth.RefreshAndGetClient(account.RefreshToken, account)
	account.AccessToken = accessToken
	return account, err
}

// ExtendsPlayerAccountTTLByAccessToken extends the TTL of the access data of the account associated by the access token.
func ExtendsPlayerAccountTTLByAccessToken(accessToken string, newDate time.Time) error {
	return tokauth.SetRemainingTime(accessToken, newDate)
}

// AuthorizeAccountByAccessToken checks that the account is authorized and returns the account found.
func AuthorizeAccountByAccessToken(accessToken string) (*models.PlayerAccount, error){
	refreshToken, err := tokauth.GetRefreshTokenFromAccessToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid or expired access token")
	}
	account := &models.PlayerAccount{
		Client: tokauth.Client{
			AccessToken: accessToken,
		},
	}
	err = c(accountCollectionName).Find(bson.M{"refreshToken": refreshToken}).One(&account)
	return account, err
}
