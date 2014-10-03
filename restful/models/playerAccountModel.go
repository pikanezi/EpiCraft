package models

import (
	"code.google.com/p/go.crypto/bcrypt"
	"fmt"
	"github.com/pikanezi/tokauth"
)

// PlayerAccount represents the account of a player.
type PlayerAccount struct {
	tokauth.Client `bson:",inline"`

	Login      string       `bson:"login" json:"login,omitempty"`
	Password   string       `bson:"password" json:"password,omitempty"`
	Characters []*Character `bson:"characters" json:"characters,omitempty"`
}

// Secure returns a secured account.
func (p *PlayerAccount) Secure() *PlayerAccount {
	p.Password = ""
	p.RefreshToken = ""
	return p
}

// WithRefreshToken returns an account with its access Token.
func (p *PlayerAccount) WithRefreshToken() *PlayerAccount {
	p.Password = ""
	return p
}

// ValidateAndCryptPassword checks if the account have every fields required and crypt the password.
func (p *PlayerAccount) ValidateAndCryptPassword() (err error) {
	if p.Login == "" || p.Password == "" {
		return fmt.Errorf("required fields are missing")
	}
	genPass, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.Password = string(genPass)
	return nil
}
