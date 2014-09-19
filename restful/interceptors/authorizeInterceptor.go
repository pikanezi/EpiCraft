package interceptors

import (
	"EpiCraft/restful/db"
	"github.com/gocarina/pi"
)

const (
	apiTokenHeader = "X-Api-Token"
)

// AuthorizeAccount checks that the request sender has the right to access these data.
func AuthorizeAccount(c *pi.RequestContext) error {
	user, err := db.AuthorizeAccountByAccessToken(c.GetHeader(apiTokenHeader))
	if err != nil {
		return pi.NewError(401, err)
	}
	c.Data["user"] = user
	return nil
}
